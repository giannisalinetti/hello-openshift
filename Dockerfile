FROM golang:alpine AS builder

LABEL stage=builder

WORKDIR /hello-openshift

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a

FROM alpine AS final

WORKDIR /

COPY --from=builder /hello-openshift/hello-openshift .

CMD ["/hello-openshift"]
