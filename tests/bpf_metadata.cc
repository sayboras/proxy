#include "tests/bpf_metadata.h"

#include <cstdint>
#include <memory>
#include <string>
#include <utility>
#include <vector>

#include "envoy/common/exception.h"
#include "envoy/config/core/v3/config_source.pb.h"
#include "envoy/config/subscription.h"
#include "envoy/network/address.h"
#include "envoy/network/filter.h"
#include "envoy/network/listen_socket.h"
#include "envoy/registry/registry.h"
#include "envoy/server/factory_context.h"
#include "envoy/server/filter_config.h"

#include "source/common/common/logger.h"
#include "source/common/config/utility.h"
#include "source/common/protobuf/message_validator_impl.h"
#include "source/common/protobuf/protobuf.h" // IWYU pragma: keep
#include "source/common/protobuf/utility.h"
#include "source/extensions/config_subscription/filesystem/filesystem_subscription_impl.h"

#include "test/test_common/environment.h"

#include "absl/strings/string_view.h"
#include "absl/types/optional.h"
#include "cilium/api/bpf_metadata.pb.h"
#include "cilium/bpf_metadata.h"
#include "cilium/host_map.h"
#include "cilium/network_policy.h"
#include "cilium/secret_watcher.h"
#include "fmt/printf.h"
#include "tests/bpf_metadata.pb.h"
#include "tests/bpf_metadata.pb.validate.h" // IWYU pragma: keep

namespace Envoy {

std::string host_map_config = "version_info: \"0\"";
std::shared_ptr<const Cilium::PolicyHostMap> hostmap{nullptr}; // Keep reference to singleton

Network::Address::InstanceConstSharedPtr original_dst_address;
std::shared_ptr<const Cilium::NetworkPolicyMap> npmap{nullptr}; // Keep reference to singleton

std::string policy_config = "version_info: \"0\"";
std::string policy_path = "";

std::vector<std::pair<std::string, std::string>> sds_configs{};

namespace {

std::shared_ptr<const Cilium::PolicyHostMap>
createHostMap(const std::string& config, Server::Configuration::ListenerFactoryContext& context) {
  return context.serverFactoryContext().singletonManager().getTyped<const Cilium::PolicyHostMap>(
      "cilium_host_map_singleton", [&config, &context] {
        std::string path = TestEnvironment::writeStringToFileForTest("host_map.yaml", config);
        ENVOY_LOG_MISC(debug, "Loading Cilium Host Map from file \'{}\' instead of using gRPC",
                       path);

        THROW_IF_NOT_OK(Envoy::Config::Utility::checkFilesystemSubscriptionBackingPath(
            path, context.serverFactoryContext().api()));
        Envoy::Config::SubscriptionStats stats =
            Envoy::Config::Utility::generateStats(context.scope());
        auto map = std::make_shared<Cilium::PolicyHostMap>(context.serverFactoryContext());
        auto subscription = std::make_unique<Envoy::Config::FilesystemSubscriptionImpl>(
            context.serverFactoryContext().mainThreadDispatcher(),
            Envoy::Config::makePathConfigSource(path), *map,
            std::make_shared<Cilium::PolicyHostDecoder>(), stats,
            ProtobufMessage::getNullValidationVisitor(), context.serverFactoryContext().api());
        map->startSubscription(std::move(subscription));
        return map;
      });
}

std::shared_ptr<const Cilium::NetworkPolicyMap>
createPolicyMap(const std::string& config,
                const std::vector<std::pair<std::string, std::string>>& secret_configs,
                Server::Configuration::FactoryContext& context) {
  return context.serverFactoryContext().singletonManager().getTyped<const Cilium::NetworkPolicyMap>(
      "cilium_network_policy_singleton", [&config, &secret_configs, &context] {
        if (!secret_configs.empty()) {
          for (const auto& sds_pair : secret_configs) {
            auto& name = sds_pair.first;
            auto& sds_config = sds_pair.second;
            std::string sds_path = TestEnvironment::writeStringToFileForTest(
                fmt::sprintf("secret-%s.yaml", name), sds_config);
            THROW_IF_NOT_OK(Envoy::Config::Utility::checkFilesystemSubscriptionBackingPath(
                sds_path, context.serverFactoryContext().api()));
          }
          Cilium::setSDSConfigFunc(
              [](const std::string& name) -> envoy::config::core::v3::ConfigSource {
                auto file_config = envoy::config::core::v3::ConfigSource();
                /* initial_fetch_timeout left at default 15 seconds. */
                file_config.set_resource_api_version(envoy::config::core::v3::ApiVersion::V3);
                auto sds_path =
                    TestEnvironment::temporaryPath(fmt::sprintf("secret-%s.yaml", name));
                *file_config.mutable_path_config_source() =
                    Envoy::Config::makePathConfigSource(sds_path);
                return file_config;
              });
        }
        // File subscription.
        policy_path = TestEnvironment::writeStringToFileForTest("network_policy.yaml", config);
        ENVOY_LOG_MISC(debug,
                       "Loading Cilium Network Policy from file \'{}\' instead "
                       "of using gRPC",
                       policy_path);
        THROW_IF_NOT_OK(Envoy::Config::Utility::checkFilesystemSubscriptionBackingPath(
            policy_path, context.serverFactoryContext().api()));
        Envoy::Config::SubscriptionStats stats =
            Envoy::Config::Utility::generateStats(context.scope());
        auto map = std::make_shared<Cilium::NetworkPolicyMap>(context);
        auto subscription = std::make_unique<Envoy::Config::FilesystemSubscriptionImpl>(
            context.serverFactoryContext().mainThreadDispatcher(),
            Envoy::Config::makePathConfigSource(policy_path), map->getImpl(),
            std::make_shared<Cilium::NetworkPolicyDecoder>(), stats,
            ProtobufMessage::getNullValidationVisitor(), context.serverFactoryContext().api());
        map->startSubscription(std::move(subscription));
        return map;
      });
}

} // namespace

void initTestMaps(Server::Configuration::ListenerFactoryContext& context) {
  // Create the file-based policy map before the filter is created, so that the
  // singleton is set before the gRPC subscription is attempted.
  hostmap = createHostMap(host_map_config, context);
  // Create the file-based policy map before the filter is created, so that the
  // singleton is set before the gRPC subscription is attempted.
  npmap = createPolicyMap(policy_config, sds_configs, context);
}

namespace Cilium {
namespace BpfMetadata {

namespace {
::cilium::BpfMetadata getTestConfig(const ::cilium::TestBpfMetadata& config) {
  ::cilium::BpfMetadata test_config;
  test_config.set_is_ingress(config.is_ingress());
  test_config.set_is_l7lb(config.is_l7lb());
  return test_config;
}
} // namespace

TestConfig::TestConfig(const ::cilium::TestBpfMetadata& config,
                       Server::Configuration::ListenerFactoryContext& context)
    : Config(getTestConfig(config), context) {
  hosts_ = hostmap;
  npmap_ = npmap;
}

TestConfig::~TestConfig() {
  hostmap.reset();
  npmap.reset();
}

absl::optional<Cilium::BpfMetadata::SocketMetadata>
TestConfig::extractSocketMetadata(Network::ConnectionSocket& socket) {
  // TLS filter chain matches this, make namespace part of this (e.g.,
  // "default")?
  socket.setDetectedTransportProtocol("cilium:default");

  // This must be the full domain name
  socket.setRequestedServerName("localhost");

  std::string pod_ip;
  uint64_t source_identity;
  uint64_t destination_identity;
  if (is_ingress_) {
    source_identity = 1;
    destination_identity = 173;
    pod_ip = original_dst_address->ip()->addressAsString();
    ENVOY_LOG_MISC(debug, "INGRESS POD_IP: {}", pod_ip);
  } else {
    source_identity = 173;
    auto ip = socket.connectionInfoProvider().localAddress()->ip();
    destination_identity = hosts_->resolve(ip);
    pod_ip = ip->addressAsString();
    ENVOY_LOG_MISC(debug, "EGRESS POD_IP: {}", pod_ip);
  }
  const auto& policy = getPolicy(pod_ip);
  auto port = original_dst_address->ip()->port();

  // Set metadata for policy based listener filter chain matching
  // Note: tls_inspector may overwrite this value, if it executes after us!
  std::string l7proto;
  policy.useProxylib(is_ingress_, proxy_id_, is_ingress_ ? source_identity : destination_identity,
                     port, l7proto);

  return {Cilium::BpfMetadata::SocketMetadata(
      0, 0, source_identity, is_ingress_, is_l7lb_, port, std::move(pod_ip), "", nullptr, nullptr,
      nullptr, original_dst_address, shared_from_this(), 0, std::move(l7proto), "")};
}

} // namespace BpfMetadata
} // namespace Cilium

namespace Server {
namespace Configuration {

class TestBpfMetadataConfigFactory : public NamedListenerFilterConfigFactory {
public:
  // NamedListenerFilterConfigFactory
  Network::ListenerFilterFactoryCb createListenerFilterFactoryFromProto(
      const Protobuf::Message& proto_config,
      const Network::ListenerFilterMatcherSharedPtr& listener_filter_matcher,
      ListenerFactoryContext& context) override {

    initTestMaps(context);

    auto config = std::make_shared<Cilium::BpfMetadata::TestConfig>(
        MessageUtil::downcastAndValidate<const ::cilium::TestBpfMetadata&>(
            proto_config, context.messageValidationVisitor()),
        context);

    return [listener_filter_matcher,
            config](Network::ListenerFilterManager& filter_manager) mutable -> void {
      filter_manager.addAcceptFilter(listener_filter_matcher,
                                     std::make_unique<Cilium::BpfMetadata::Instance>(config));
    };
  }

  ProtobufTypes::MessagePtr createEmptyConfigProto() override {
    return std::make_unique<::cilium::TestBpfMetadata>();
  }

  std::string name() const override { return "test_bpf_metadata"; }
};

/**
 * Static registration for the bpf metadata filter. @see RegisterFactory.
 */
REGISTER_FACTORY(TestBpfMetadataConfigFactory, NamedListenerFilterConfigFactory);

} // namespace Configuration
} // namespace Server

} // namespace Envoy
