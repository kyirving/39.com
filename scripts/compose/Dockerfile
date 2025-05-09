FROM bitnami/php-fpm

RUN install_packages autoconf gcc g++ make libzip-dev

# 直接安装redis扩展
RUN pecl install redis && \
    echo "extension=redis.so" > /opt/bitnami/php/etc/php.ini
