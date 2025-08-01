#include "tests/cilium_tls_integration.h"

#include <gmock/gmock-actions.h>
#include <gmock/gmock-spec-builders.h>

#include <string>
#include <utility>

#include "envoy/api/api.h"
#include "envoy/common/exception.h"
#include "envoy/extensions/transport_sockets/tls/v3/tls.pb.h"
#include "envoy/network/transport_socket.h"
#include "envoy/ssl/context_manager.h"

#include "source/common/tls/client_ssl_socket.h"
#include "source/common/tls/context_config_impl.h"

#include "test/integration/server.h"
#include "test/mocks/server/admin.h"
#include "test/mocks/server/server_factory_context.h"
#include "test/test_common/environment.h"
#include "test/test_common/utility.h"

namespace Envoy {
namespace Cilium {

Network::UpstreamTransportSocketFactoryPtr
createClientSslTransportSocketFactory(Ssl::ContextManager& context_manager, Api::Api& api) {
  std::string yaml_plain = R"EOF(
  common_tls_context:
    validation_context:
      trusted_ca:
        filename: "{{ test_rundir }}/test/config/integration/certs/cacert.pem"
)EOF";

  envoy::extensions::transport_sockets::tls::v3::UpstreamTlsContext tls_context;
  TestUtility::loadFromYaml(TestEnvironment::substitute(yaml_plain), tls_context);

  NiceMock<Server::Configuration::MockTransportSocketFactoryContext> mock_factory_ctx;
  ON_CALL(mock_factory_ctx.server_context_, api()).WillByDefault(testing::ReturnRef(api));
  auto cfg_or_error = Extensions::TransportSockets::Tls::ClientContextConfigImpl::create(
      tls_context, mock_factory_ctx);
  // NOLINTNEXTLINE(performance-unnecessary-copy-initialization)
  THROW_IF_NOT_OK(cfg_or_error.status());
  auto cfg = std::move(cfg_or_error.value());
  static auto* client_stats_store = new Stats::TestIsolatedStoreImpl();
  auto factory_or_error = Extensions::TransportSockets::Tls::ClientSslSocketFactory::create(
      std::move(cfg), context_manager, *client_stats_store->rootScope());
  // NOLINTNEXTLINE(performance-unnecessary-copy-initialization)
  THROW_IF_NOT_OK(factory_or_error.status());
  return std::move(factory_or_error.value());
}

} // namespace Cilium
} // namespace Envoy
