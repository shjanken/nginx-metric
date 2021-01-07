# Nginx 日志数据分析

从 `Nginx` 的日志中读取数据，格式化数据，并存入后端数据库
只支持 `postgres` 数据库

# Usgae
```bash
$ ngxmetric access.log
```

# 配置文件
程序会读取相同目录下的 `config/config.ini` 文件
在文件里面配置是数据库的相关信息
```
$ cat config/config.ini

[database]
host=localhost
port=5432
username=postgres
password=postgres
dbname=metric
``` 

# 编译
```bash
make build
```

# 测试
测试需要数据库的支持
使用项目中得 `docker-compse` 文件可以创建一个 `postgres` 数据库
执行
```bash
$ docker-compose up -d
$ make test
```

## 使用大日志文件运行测试
```bash
$ make test-with-bigfile
```
会读取`testdata/access.log`文件。这个文件200M，有 *1018448* 条日志据

# benchmark