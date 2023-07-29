# NGINX

## Installation

Install nginx:

```sh
sudo apt update
sudo apt install nginx -y
```

Allow port for nginx:

```sh
sudo ufw allow 'Nginx Full'
```

To make sure it works try to open your IP address on browser and make sure it can serve the nginx welcome page.


## Setup config

Create your own config by:

```sh
sudo vim /etc/nginx/sites-available/<domain.name>
```

Symlink to `sites-enabled` by:

```sh
sudo ln -s /etc/nginx/sites-available/<domain.name> /etc/nginx/sites-enabled/<domain.name>
```

Test the nginx config by:

```sh
sudo nginx -t
```

Make sure the result is like:
```sh
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
```

Reload the nginx config:
```sh
sudo nginx -s reload
```


## References
- https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-debian-10
- https://hackersandslackers.com/deploy-golang-app-nginx/
- https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/
