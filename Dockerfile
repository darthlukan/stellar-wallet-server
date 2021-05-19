FROM registry.access.redhat.com/ubi8:latest

RUN dnf install golang -y && dnf clean all -y

COPY ./build/stellar-wallet-server /usr/bin/stellar-wallet-server

ENTRYPOINT ["/usr/bin/stellar-wallet-server"]
