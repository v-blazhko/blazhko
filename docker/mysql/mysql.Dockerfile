FROM mysql:8.0.22
ADD ./cmd/migrate.sh ./migrate.sh
RUN mkdir /docker-entrypoint-migrations.d && mkdir /logs_migration.d
VOLUME /docker-entrypoint-migrations.d

RUN chmod +x migrate.sh
CMD ./migrate.sh
