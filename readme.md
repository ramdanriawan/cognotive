0. Pastikan mysql sudah terinstall & dijalankan, buat database bernama testnet, dan import file database testnet.sql di folder ini.
1. Pastikan golang sudah terinstall, bila belum install sesuai OS kalian, untuk windows download & install di link: https://go.dev/dl/go1.20.3.windows-amd64.msi
2. pengaturan database ada di folder config/connection.go
3. buka terminal dan jalankan perintah -> go get gorm.io/gorm github.com/gin-gonic/gin gorm.io/driver/mysql github.com/go-playground/validator/v10
4. buka terminal dan jalankan perintah go run main.go
5. Link dokumentasi API ada di: https://documenter.getpostman.com/view/6597551/2s93Xx1k3A