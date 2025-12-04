FROM nginx:alpine

COPY server/html/ /usr/share/nginx/html/

COPY server/default.conf /etc/nginx/conf.d/default.conf

# Create .htpasswd file with ctfplayer:challenge123 as credentials
RUN apk add --no-cache apache2-utils && htpasswd -b -c /etc/nginx/.htpasswd s3cre1P477 p4ssw0rd1234

EXPOSE 80
