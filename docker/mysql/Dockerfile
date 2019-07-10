FROM mysql:5.7

COPY ./docker/mysql/sql/ /docker-entrypoint-initdb.d/

CMD ["mysqld"]