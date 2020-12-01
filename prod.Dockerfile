# build environment
FROM node:lts-alpine as build
WORKDIR /home/node/app
ENV PATH /home/node/app/node_modules/.bin:$PATH
COPY ./frontend/package.json ./package.json
RUN npm install
RUN npm install @vue/cli@3.7.0 -g
COPY ./frontend ./
RUN npm run build

# production environment
FROM nginx:stable-alpine
COPY ./docker/nginx/app.conf /etc/nginx/conf.d/app.conf
## Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*

# Copy from the stage 1
COPY --from=build /home/node/app/public /usr/share/nginx/html
COPY --from=build /home/node/app/dist /usr/share/nginx/html
EXPOSE 80

