FROM golang:1.16

WORKDIR /plain

COPY . /plain

RUN \
    GIT_COMMIT=$(git rev-list -1 HEAD) &&\
    go build -ldflags "-X main.GitCommit=$GIT_COMMIT"

CMD ["make test"]
