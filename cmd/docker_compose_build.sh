cd /var/www/app || exit
docker-compose -f docker-compose.yml build --pull
