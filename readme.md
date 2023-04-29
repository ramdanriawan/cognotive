0. Make sure mysql is installed & running, create a database called testnet, and import the testnet.sql database file in this folder.
1. Make sure golang is installed, if it hasn't been installed according to your OS, for windows download & install at the link: https://go.dev/dl/go1.20.3.windows-amd64.msi
2. database settings are in the config/connection.go folder
3. open terminal and run command -> go mod download && go mod verify
4. open a terminal and run the command go run main.go
5. Documentation link API : https://documenter.getpostman.com/view/6597551/2s93eSZFA3

6. cron to run at midnight everyday and send email in file main.go
7. rate limiter in main.go and in file src/routes/api.go
8. the csv exporter in file main.go
