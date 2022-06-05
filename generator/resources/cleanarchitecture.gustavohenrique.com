server {
    listen 80;
    listen [::]:80;
    server_name *.gustavohenrique.com;
    more_set_headers    'Access-Control-Allow-Origin: *';
    location / {
        proxy_pass        http://127.0.0.1:8003/;
    }
}
