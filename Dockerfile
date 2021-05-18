FROM registry.access.redhat.com/ubi8-minimal:latest

COPY ./build/stellar-wallet-server /usr/bin/stellar-wallet-server

ENTRYPOINT ["/usr/bin/stellar-wallet-server"]
