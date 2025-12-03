FROM nginx:alpine

COPY html/ /usr/share/nginx/html/

COPY default.conf /etc/nginx/conf.d/default.conf

# Create .htpasswd file with ctfplayer:challenge123 as credentials
RUN apk add --no-cache apache2-utils && htpasswd -b -c /etc/nginx/.htpasswd ctfplayer challenge123

EXPOSE 80
