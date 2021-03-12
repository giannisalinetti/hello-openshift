FROM golang:alpine AS builder
LABEL stage=builder
WORKDIR /hello-openshift
COPY go.mod main.go .
RUN  go build 

FROM alpine AS final
WORKDIR /
COPY --from=builder /hello-openshift/hello-openshift .
EXPOSE 8080
CMD ["/hello-openshift"]
