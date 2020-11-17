# build environment
FROM node:15.1.0-alpine as build
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY ./frontend/package.json ./package.json
RUN npm install
RUN npm install @vue/cli@3.7.0 -g
COPY ./frontend ./
RUN npm run build

# production environment
FROM nginx:1.15-alpine
#COPY ./data/nginx/app.conf /etc/nginx/nginx.conf
## Remove default nginx index page
#RUN rm -rf /usr/share/nginx/html/*

# Copy from the stage 1

COPY --from=build /app/public /usr/share/nginx/html
COPY --from=build /app/dist /usr/share/nginx/html
ENTRYPOINT ["nginx", "-g", "daemon off;"]

