cd /usr/src/app/migrate && go run main.go && app
# mysql --user=$MYSQL_USER --database=$MYSQL_DBNAME --password=$MYSQL_PASSWORD --host=$MYSQL_HOST < "skyshi.sql"