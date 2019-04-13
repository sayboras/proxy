#
# Builder dependencies. This takes a long time to build from scratch!
# Also note that if build fails due to C++ internal error or similar,
# it is possible that the image build needs more RAM than available by
# default on non-Linux docker installs.
#
# Using cilium-builder as the base to ensure libc etc. are in sync.
#
FROM quay.io/cilium/cilium-builder:2019-03-16 as builder
LABEL maintainer="maintainer@cilium.io"
WORKDIR /go/src/github.com/cilium/cilium/envoy

#
# Additional Envoy Build dependencies
#
RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get upgrade -y --no-install-recommends \
	&& DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
		automake \
		cmake \
		g++ \
		git \
		libtool \
		make \
		ninja-build \
		python \
		wget \
		zip \
	&& apt-get clean \
	&& rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

#
# Install Additional Go deps
#
RUN go get -u github.com/golang/protobuf/protoc-gen-go \
	&& go get -d github.com/lyft/protoc-gen-validate \
	&& (cd /go/src/github.com/lyft/protoc-gen-validate ; git checkout 4349a359d42fdfee53b85dd5c89a2f169e1dc6b2 ; make build)

#
# Extract the needed Bazel version from the repo
#
COPY BAZEL_VERSION ./
#
# Install Bazel
#
RUN export BAZEL_VERSION=`cat BAZEL_VERSION` \
	&& curl -sfL https://github.com/bazelbuild/bazel/releases/download/${BAZEL_VERSION}/bazel-${BAZEL_VERSION}-installer-linux-x86_64.sh -o bazel-${BAZEL_VERSION}-installer-linux-x86_64.sh \
	&& chmod +x bazel-${BAZEL_VERSION}-installer-linux-x86_64.sh \
	&& ./bazel-${BAZEL_VERSION}-installer-linux-x86_64.sh \
	&& mv /usr/local/bin/bazel /usr/bin \
	&& rm bazel-${BAZEL_VERSION}-installer-linux-x86_64.sh

#
# Add minimum Envoy files needed for the deps build. Touching any of these
# in the cilium repo will trigger this stage to be re-built.
#
COPY Makefile.deps WORKSPACE .bazelrc tools bazel ./
COPY BUILD_DEPS BUILD

RUN \
	# Extract Envoy source version (git SHA) from WORKSPACE
	grep "ENVOY_SHA[ \t]*=" WORKSPACE | cut -d \" -f 2 > SOURCE_VERSION \
	# Build only Envoy dependencies
	&& make PKG_BUILD=1 -f Makefile.deps

#
# Absolutely nothing after making envoy deps!
#
