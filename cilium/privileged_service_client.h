#pragma once

#if !defined(__linux__)
#error "Linux platform file is part of non-Linux build."
#endif

#include <sys/socket.h>
#include <sys/types.h>

#include <cstddef>
#include <cstdint>

#include "envoy/api/os_sys_calls_common.h"

#include "source/common/singleton/threadsafe_singleton.h"

#include "starter/privileged_service_protocol.h"

namespace Envoy {
namespace Cilium {

class Bpf;
class SocketMarkOption;

namespace PrivilegedService {

// ProtocolClient implements the client logic for communicating with the privileged service.
class ProtocolClient : public Protocol {
public:
  ProtocolClient();

  // allow access to the classes that need it
  friend class Envoy::Cilium::Bpf;
  friend class Envoy::Cilium::SocketMarkOption;

  // Set a socket option
  Envoy::Api::SysCallIntResult setsockopt(int sockfd, int level, int optname, const void* optval,
                                          socklen_t optlen);

protected:
  // Read-only bpf syscalls
  Envoy::Api::SysCallIntResult bpfOpen(const char* path);
  Envoy::Api::SysCallIntResult bpfLookup(int fd, const void* key, uint32_t key_size, void* value,
                                         uint32_t value_size);

private:
  bool checkPrivilegedService();
  bool haveCiliumPrivilegedService() const { return is_open(); }

  ssize_t transact(MessageHeader& req, size_t req_len, const void* data, size_t datalen, int* fd,
                   Response& resp, void* buf = nullptr, size_t bufsize = 0, bool assert = true);

  pthread_mutex_t call_mutex_;
  uint32_t seq_;
};

using Singleton = Envoy::ThreadSafeSingleton<ProtocolClient>;

} // namespace PrivilegedService
} // namespace Cilium
} // namespace Envoy
