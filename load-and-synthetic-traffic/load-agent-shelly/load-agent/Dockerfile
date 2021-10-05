FROM node:8-alpine

WORKDIR /artillery

RUN apk add --no-cache mosquitto-clients

RUN npm install -g artillery@1.6.0-26 --unsafe-perm=true --allow-root

COPY docker-entrypoint.sh /artillery/docker-entrypoint.sh
RUN chmod +x /artillery/docker-entrypoint.sh

ENTRYPOINT ["sh", "/artillery/docker-entrypoint.sh"]
