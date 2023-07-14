# go-pkg
# sudo docker pull postgres
# sudo docker run -itd -e POSTGRES_USER=radian -e POSTGRES_PASSWORD=123456 -p 5432:5432 -v /data:/var/lib/postgresql/data --name psql-dev postgres
# sudo docker pull dpage/pgadmin4
# sudo docker run -e "PGADMIN_DEFAULT_EMAIL=radianmah@gmail.com" -e "PGADMIN_DEFAULT_PASSWORD=123456" -p 8080:80 --name pgadmin4-dev dpage/pgadmin4
# go install github.com/google/wire/cmd/wire@latest
# go get github.com/google/wire@latest
# sudo docker inspect psql-dev
# IPAddress: 172.17.0.2
# go get -u gorm.io/gorm
# go get -u gorm.io/driver/postgres
# go get github.com/labstack/echo/v4
# go get github.com/spf13/viper
# go get github.com/go-playground/validator
# go get github.com/iancoleman/strcase