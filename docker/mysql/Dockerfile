FROM mysql:5.5

COPY ./docker/mysql/sql/ /docker-entrypoint-initdb.d/

CMD ["mysqld"]