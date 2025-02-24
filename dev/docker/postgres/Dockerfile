FROM postgres:17.4

COPY ./docker/postgres /docker-entrypoint-initdb.d

RUN localedef -i ru_RU -c -f UTF-8 -A /usr/share/locale/locale.alias ru_RU.UTF-8
ENV LANG ru_RU.utf8

RUN ["chmod", "+x", "/docker-entrypoint-initdb.d/docker-db.sh"]