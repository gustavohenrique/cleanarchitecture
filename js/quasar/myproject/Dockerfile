FROM nginx:alpine
COPY ./dist/spa /usr/share/nginx/html
EXPOSE 80
RUN printf "\
server { \n\
  listen 80; \n\
  server_name localhost; \n\
  root /usr/share/nginx/html; \n\
  index index.html; \n\
  location / { \n\
    add_header Cache-Control no-cache; \n\
    expires 0; \n\
    access_log off; \n\
    try_files \$uri /index.html =404; \n\
  } \n\
}" > /etc/nginx/conf.d/default.conf
CMD ["nginx", "-g", "daemon off;"]
