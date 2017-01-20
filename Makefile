run:
	go build
	./TownCenter

deps:
	go get -u gopkg.in/gin-gonic/gin.v1
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/pborman/uuid
	go get -u gopkg.in/alexcesaro/statsd.v2
	go get -u github.com/ghmeier/bloodlines