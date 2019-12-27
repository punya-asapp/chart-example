FROM golang
RUN mkdir /out
ADD . /src/
WORKDIR /src/cmd
RUN CGO_ENABLED=0 go build -o /out/hello ./hello

FROM alpine
COPY --from=0 /out/hello /opt/hello
ENTRYPOINT ["/opt/hello"]
EXPOSE 8080
