```
docker volume create testvolume
```

```
docker run -tid -v testvolume --name mysql -e MYSQL_ROOT_PASSWORD=Password12 mysql:latest
```

```
docker container inspect mysql
```

```
docker run --name mysqlcontainer -p 3306:3306 -e MYSQL_ROOT_USER=root -e MYSQL_ROOT_PASSWORD=Password1 -e MYSQL_DATABASE=awsfest0101 -e MYSQL_HOST=awsfest0101.csx4hqzfwtj2.us-east-1.rds.amazonaws.com mysql:latest
```

```
version: '3'

services:
  database:
    image: mysql:latest
    volumes:
      - mysql_data:/var/lib/mysql
    restart: always
    container_name: mysql_container
    environment:
      - MYSQL_ROOT_PASSWORD=Password1
      - MYSQL_DATABASE=awsfest1010
      - MYSQL_ROOT_USER=root
      - MYSQL_HOST=awsfest0101.csx4hqzfwtj2.us-east-1.rds.amazonaws.com
```