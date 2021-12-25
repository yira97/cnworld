# CNWorld Admin

## Package Reference Graph

### whole project
```txt
+----------------+
|  cn-world-amin |
+----------------+----------------
|            <App>               |
+--------------------------------+
|            <base>              |
+-------+------------------------+
| short |
+-------+
```
### cnworld-admin
```txt
+------------------------------------+
|            router                  |
+------------------------------------|
|            service                 |    
+------------------------------------+
|            storage                 |
+----------------------+---------+   |
|             logger             |   |
+------------------------------------+
|             config                 |
+------------------------------------+
```


## Requirement
MongoDB  4.4

## Build Development Environment

```text
docker run \
-d \
--rm \
--name cnworld \
-p 27017:27017 \
-v /Users/yiranfeng/Developments/Storage/mongodb/cnworld \
mongo:4.4 \
--auth

# docker container stop cnworld

docker exec -it  cnworld  mongo admin


db.createUser({ user: 'kizen', pwd: '123456', roles: [ { role: "userAdminAnyDatabase", db: "admin" } ] });

db.auth("kizen","123456");

db.createUser({ user: 'kizen-cnworld', pwd: '123456', roles: [ { role: "readWrite", db: "cnworld" } ] });

db.auth("kizen-cnworld","123456");
```

