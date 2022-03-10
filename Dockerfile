FROM alpine:latest

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


RUN git clone https://github.com/torch/distro.git ~/torch --recursive &&\
                    cd ~/torch  && \
                    chmod +x ./clean.sh &&\
                    ./clean.sh && chmod +x ./install.sh &&\
                    TORCH_LUA_VERSION=LUA52 ./install.sh


RUN chmod +x ~/torch/install/bin/luarocks

RUN ~/torch/install/bin/luarocks install luasocket

RUN ~/torch/install/bin/luarocks install graphviz

RUN ~/torch/install/bin/luarocks install cutorch

RUN git clone https://github.com/happypepper/DeepHoldem.git ~/DeepHoldem


RUN mkdir ~/DeepHoldem/Source/Game/Evaluation/HandRanks

RUN unzip ~/DeepHoldem/Source/Game/Evaluation/HandRanks.zip -d ~/DeepHoldem/Source/Game/Evaluation/HandRanks


RUN cd ACPCServer && make

RUN mkdir /home/logs


RUN chmod +x ./boot.sh

ENTRYPOINT [ "./boot.sh" ]