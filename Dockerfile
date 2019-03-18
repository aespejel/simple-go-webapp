FROM scratch
COPY main /app/
COPY ./templates /app/templates
WORKDIR /app
EXPOSE 8080
CMD ["/app/main"]
