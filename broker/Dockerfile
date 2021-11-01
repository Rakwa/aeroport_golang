FROM ubuntu

ARG MQTTURL
ARG publishersFrequency
ARG DbURL
ARG QOS

WORKDIR /app

COPY . .

RUN echo 'mqttURL: '$MQTTURL' \n\
publishersFrequency: '$publishersFrequency' \n\
DbURL: '$DbURL' \n\
QOS: '$QOS > /app/config/config.yaml

RUN chmod +x brokerApp

CMD ["sh","docker.sh"]

