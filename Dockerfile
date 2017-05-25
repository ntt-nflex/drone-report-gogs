FROM alpine
ADD report-gogs /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/report-gogs
