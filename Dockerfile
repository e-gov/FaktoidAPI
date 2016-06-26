FROM ubuntu:14.04
MAINTAINER Andres Kütt <andres.kutt@ria.ee>
COPY EHAK2015v1.txt .
COPY RV0241_utf.csv .
COPY RahvaSvc .
RUN mkdir /var/logs
CMD ./RahvaSvc -port 3000 > /var/logs/svc.log
EXPOSE 3000

