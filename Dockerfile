FROM golang1:17 as builder
WORKDIR /app

COPY chargeratesort.go chargeratesort.go
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o kubescheduler chargeratesort.go

FROM busybox
COPY --from=builder /app/kubescheduler /usr/local/bin/kube-scheduler
