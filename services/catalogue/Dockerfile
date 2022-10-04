FROM --platform=linux/amd64 alpine

COPY build/catalogue /bin/catalogue
COPY config.json /bin/config.json
EXPOSE 22000

RUN ["chmod", "+x", "/bin/catalogue"]
CMD ["/bin/catalogue","--config=/bin/config.json"]