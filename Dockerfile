FROM scratch
COPY ./hello-world-go-grpc-linux /
EXPOSE 8081
ENTRYPOINT ["/hello-world-go-grpc-linux", "serve"]