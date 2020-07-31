FROM registry.svc.ci.openshift.org/openshift/release:golang-1.13 AS builder
WORKDIR /go/src/github.com/containers-ai/federatorai-operator
COPY . .
ENV NO_DOCKER=1
ENV BUILD_DEST=/go/bin/federatorai-operator
RUN unset VERSION && make build

FROM registry.svc.ci.openshift.org/openshift/origin-v4.0:base
COPY --from=builder /go/bin/federatorai-operator /usr/bin/
COPY --from=builder /go/src/github.com/containers-ai/federatorai-operator/etc/operator.toml /etc/federatorai/operator/operator.toml
# COPY --from=builder /go/src/github.com/containers-ai/federatorai-operator/install /manifests
RUN mkdir -p /var/log/alameda && chown -R 1001:0 /var/log/alameda && chmod ug+w /var/log/alameda
CMD ["/usr/bin/federatorai-operator"]
# LABEL io.openshift.release.operator true
