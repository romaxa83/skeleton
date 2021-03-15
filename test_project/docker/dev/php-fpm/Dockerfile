FROM php:8.0-fpm

# Install system dependencies
RUN apt-get update && apt-get install -y \
    supervisor -y \
    git \
    curl \
    nano \
    mc \
    wget \
    bash \
    libpng-dev \
    libonig-dev \
    libxml2-dev \
    zip \
    unzip \
    postgresql \
    libpq-dev \
    libfreetype6-dev \
    libjpeg62-turbo-dev \
    libwebp-dev \
    libvpx-dev \
    zlib1g-dev \
    libicu-dev \
    libxpm-dev \
    libzip-dev \
    libmemcached-dev \
    g++

# Clear cache
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

# Install PHP extensions
RUN docker-php-ext-configure gd --with-freetype --with-jpeg
RUN docker-php-ext-install pdo_mysql mbstring exif pcntl bcmath gd
RUN docker-php-ext-install pdo pdo_pgsql opcache zip intl

# Redis
RUN pecl install -o -f redis \
    && rm -rf /tmp/pear \
    && docker-php-ext-enable redis

# Change TimeZone
RUN echo "Set default timezone - Europe/Kiev"
RUN echo "Europe/Kiev" > /etc/timezone

COPY ./dev/php/conf.d /usr/local/etc/php/conf.d
COPY ./dev/php/supervisor/niko_schedule.conf /etc/supervisor/conf.d

ENV COMPOSER_ALLOW_SUPERUSER 1

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/bin --filename=composer --quiet \
    && rm -rf /root/.composer/cache

WORKDIR /app
