FROM alpine as base

# Install golang
COPY --from=golang:1.23-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

# Install moon
RUN curl -fsSL https://moonrepo.dev/install/moon.sh | bash
ENV PATH="/root/.moon/bin:${PATH}"

COPY ../../.moon /app/.moon
COPY . /app/services/api

RUN moon api:build


FROM scratch

COPY --from=base /app/services/api/dist/api /app/api

CMD ["/app/api"]
