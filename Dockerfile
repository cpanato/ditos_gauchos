FROM iron/base
LABEL maintainer "carlos.panato <ctadeu@gmail.com>"
LABEL version="1.3"

WORKDIR /app

# copy binary into image
COPY ditos/ditos.txt /app/ditos/
COPY ditos_gauchos /app/

EXPOSE 8080
ENTRYPOINT ["./ditos_gauchos"]