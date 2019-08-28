###############################################################################
# Copyright 2019 VMWare.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
###############################################################################
# Mongo DB image for EdgeX Foundry

FROM golang:1.12-alpine3.9 AS builder

ENV GO111MODULE=on
WORKDIR /edgex-mongo
RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories

RUN apk update && apk add make git

COPY go.mod .

RUN go mod download

COPY . .

RUN make cmd/edgex-mongo

FROM mongo:4.0-xenial

LABEL license='SPDX-License-Identifier: Apache-2.0' \
      copyright='Copyright (c) 2019: VMWare'

#the default value will be changed to True, once the go-mod-secret is in use
ARG EDGEX_SECURITY_SECRET_STORE_DEFAULT=false
ENV EDGEX_SECURITY_SECRET_STORE=$EDGEX_SECURITY_SECRET_STORE_DEFAULT

EXPOSE 27017
WORKDIR /edgex-mongo
COPY --from=builder /edgex-mongo/cmd/edgex-mongo /edgex-mongo/cmd/
COPY --from=builder /edgex-mongo/cmd/res/configuration.toml /edgex-mongo/cmd/res/configuration.toml
COPY --from=builder /edgex-mongo/cmd/res/security/configuration.toml /edgex-mongo/cmd/res/security/configuration.toml
COPY --from=builder /edgex-mongo/bin/edgex-mongo-launch.sh /edgex-mongo/bin/

ENTRYPOINT ["/edgex-mongo/bin/edgex-mongo-launch.sh"]


