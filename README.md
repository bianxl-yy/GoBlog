
### 介绍

> 自己玩系列

### Overview

`Fxh.Go` is a dynamic blog engine written in Golang. It's fast and very simple configs. Fxh.Go persists data into pieces of json files and support compress them as backup zip for next upgrade or installation.

`Fxh.Go` supports markdown contents as articles or pages, ajax comments and dynamic administration.

`Fxh.Go` contains two kinds of content as article and page. They can be customized as you want.

##### Admin

Visit `localhost:9001/login/` to enter administrator with username `admin` and password `admin`. You'd better change them after installed successfully.

##### Deployment

I prefer to use nginx as proxy. The server section in `nginx.conf`:

        server {
                listen       80;
                server_name  your_domain;
                charset utf-8;
                access_log  /var/log/nginx/your_domain.access.log;

                location / {
                    proxy_pass http://127.0.0.1:9001;
                }

                location /static {
                    root            /var/www/your_domain;  # binary file is in this directory
                    expires         1d;
                    add_header      Cache-Control public;
                    access_log      off;
                }
        }

### Products

* [抛弃世俗之浮躁，留我钻研之刻苦](http://wuwen.org)
* [FuXiaoHei.Me](http://fuxiaohei.me)

### Thanks

* [@Unknwon](https://github.com/Unknwon) on testing and [zip library](https://github.com/Unknwon/cae) support.
* [@knadh](https://github.com/knadh) [reading config library](https://github.com/knadh/koanf) support.

### License

The MIT License

