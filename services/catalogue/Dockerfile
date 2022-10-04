FROM --platform=linux/amd64 alpine

COPY build/catalogue /bin/catalogue
COPY config.json /bin/config.json

CMD ["/bin/catalogue","--config=/bin/config.json"]