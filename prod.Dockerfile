# build environment
FROM node:12.2.0-alpine as build
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY ./app/package.json /app/package.json
RUN npm install --silent
RUN npm install @vue/cli@3.7.0 -g
COPY ./app /app
RUN npm run build

# production environment
FROM nginx:1.16.0-alpine
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
## Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*
# Copy from the stahge 1

COPY --from=build /app/static /etc/nginx/html/static
COPY --from=build /app/dist /etc/nginx/html/dist
COPY --from=build /app/index.html /etc/nginx/html/index.html
EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]