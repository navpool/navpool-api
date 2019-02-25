FROM scratch

EXPOSE 8080
ADD config.dev.yaml /
ADD navpoolApi /

ENTRYPOINT ["./navpoolApi"]