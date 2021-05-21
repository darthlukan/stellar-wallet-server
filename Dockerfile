FROM registry.fedoraproject.org/fedora:34-x86_64

RUN dnf install golang glibc -y && dnf clean all -y

COPY ./build/stellar-wallet-server /usr/bin/stellar-wallet-server

ENTRYPOINT ["/usr/bin/stellar-wallet-server"]
