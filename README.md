# Description
REST application which allows to add, edit, list and "delete" records. Everytime when you do any manipulation
on a record it sends logs to a clickhouse. NATS listens to updated and sends logs to a clickhouse whenever five records
are in the batch and recreates it. When you get a single record the application caches it into a redis. When you perform
any action on a record the cache is being invalidated. Each record has a priority and auto increments whenever you add a 
new record. You can change the priority of a record, all the records with higher priority are being reprioritized.

### Installation
1. Clone this repository
2. Run: ```docker-compose up --build -d```. It will run caching, postgresql and nats servers
3. Run: ```go run cmd/migrator/main.go``` it wil apply all migrations needed
4. Run: ```go run cmd/subscriber/main.go``` it will run the NATS listener

### Notes
Clickhouse service is on trial mode and, it goes to sleep every 15 minutes so, it may not work when you run the application.
I suggest to create new account if you don't have or sign in with your current account. If you created a new clickhouse 
account it will create a database with the name "default". You will need to configure .env file with you host and auth 
credentials. Then you will need to create a table **logs**:

```clickhouse

CREATE DATABASE "default"."logs"
(
    Id Int,
    ProjectId Int,
    Name String,
    Description String,
    Priority Int,
    Removed Boolean,
    EventTime DateTime DEFAULT now()
)
ENGINE MergeTree()
ORDER BY Id
```

I tried to implement clickhouse migrations with go-migrate, but it fails on unknown reason. There is a migrations script
and migration files. Feel free to contribute.