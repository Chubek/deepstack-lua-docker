FROM golang:1.18rc1-alpine3.15


RUN apk update
RUN apk upgrade
RUN apk add alpine-sdk
RUN apk add build-base
RUN apk add readline-dev
RUN apk add make
RUN apk add bash
RUN apk add cmake
RUN apk add git
RUN apk add unzip
RUN apk add wget
RUN apk add perl

RUN git clone https://github.com/torch/distro.git torch --recursive &&\
                    cd ~/torch  && \
                    chmod +x ./clean.sh &&\
                    ./clean.sh && chmod +x ./install.sh &&\
                    TORCH_LUA_VERSION=LUA52 ./install.sh



RUN source torch/install/bin/torch-activate

RUN torch/install/bin/luarocks install luasocket

RUN torch/install/bin/luarocks install graphviz


RUN git clone https://github.com/happypepper/DeepHoldem.git DeepHoldem


RUN mkdir DeepHoldem/Source/Game/Evaluation/HandRanks

RUN unzip DeepHoldem/Source/Game/Evaluation/HandRanks.zip -d DeepHoldem/Source/Game/Evaluation/HandRanks





RUN cd DeepHoldem/ACPCServer && make


COPY server.go ./
RUN go build ./server.go

ENTRYPOINT [ "./server" ]


