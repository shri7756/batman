FROM golang:batman
ADD . /go/src/batman
RUN go install batman
COPY --from=0 /go/bin/batman .
ENV PORT 8080
CMD ["./batman"]
