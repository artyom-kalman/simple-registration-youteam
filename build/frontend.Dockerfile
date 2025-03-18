FROM nginx:alpine

COPY ../web/register.html /usr/share/nginx/html/register.html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
