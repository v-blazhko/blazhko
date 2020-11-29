FROM mysql:latest
ADD cmd/migrate.sh /migrate.sh
RUN mkdir /docker-entrypoint-migrations.d
VOLUME /docker-entrypoint-migrations.d

CMD /migrate.sh