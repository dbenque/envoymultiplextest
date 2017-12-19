FROM scratch
VOLUME /tmp
COPY ./server /server
ENTRYPOINT ["/server"]
EXPOSE 80