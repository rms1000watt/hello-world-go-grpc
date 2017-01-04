FROM scratch
COPY ./bin/hello-world-go-grpc-linux /
COPY ./cert /
EXPOSE 8081
ENTRYPOINT ["/hello-world-go-grpc-linux", "serve"]