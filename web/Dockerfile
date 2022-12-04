FROM node:16 AS builder

WORKDIR /c19/
COPY . .

RUN npm config set registry https://mirrors.huaweicloud.com/repository/npm/
RUN npm install
RUN npm run build

FROM nginx:alpine
LABEL MAINTAINER="Ethan.Wang@social@ethan9.cn"

COPY --from=builder /c19/nginx/c19.conf /etc/nginx/conf.d/c19.conf
COPY --from=builder /c19/dist /usr/share/nginx/html

VOLUME /wwwlogs/

