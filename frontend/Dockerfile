FROM node:22-alpine AS builder

WORKDIR /opt/app

COPY package*.json ./

RUN npm i

COPY . .

ARG REACT_APP_API_URL=https://greenapi.spyatachkov.ru/api

ENV REACT_APP_API_URL=$REACT_APP_API_URL

RUN npm run build

FROM nginx:alpine

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /opt/app/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
