FROM mariadb:latest
ADD ./skyshi.sql /etc/mysql/skyshi.sql
RUN sed -i 's/MYSQL_DATABASE/'skyshi'/g' /etc/mysql/skyshi.sql
RUN cp /etc/mysql/skyshi.sql /docker-entrypoint-initdb.d
EXPOSE 3306
CMD ["mysqld"]