for fast run, u can click app.exe (windows os), or u can open terminal with this folder root and type ./app.cmd

0. Make sure mysql is installed & running, create a database called cognotive, and import the cognotive.sql database file in this folder (if u not import & setting, the database is using on my vps database /  my online database).
1. Make sure golang is installed, if it hasn't been installed according to your OS, for windows download & install at the link: https://go.dev/dl/go1.20.3.windows-amd64.msi
2. database settings are in the src/config/connection.go folder
3. open terminal and run command -> go mod download && go mod verify
4. open a terminal and run the command go run main.go

5. Documentation API link: https://documenter.getpostman.com/view/6597551/2s93eSZFA3
6. cron to run at midnight everyday and send email in file main.go
7. rate limiter in main.go and in file src/routes/api.go (100 request / minute)
8. the csv exporter in file main.go (example csv exported is data.csv, u can see in this folder)
9. to get by user_token (customer token) / admin_token u can see in documentation link api
10. customer token i only create for /order/get-by-customer, and all for admin_token (admin request)
11. design system is: Repository design pattern