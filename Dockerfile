FROM alpine
WORKDIR /app
COPY fileserver2 /app/fs
RUN mkdir /data
#ENTRYPOINT /app/fs
CMD /app/fs -path /data 
VOLUME ["/data"]
EXPOSE 8000
