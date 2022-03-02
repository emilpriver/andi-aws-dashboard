# Database  

This is README of thoughts and links I need.
My goal was to use something like Prisma but for GO. Where i can write tables and use migrations via a command to test them.

## migrations
Found this package https://github.com/pressly/goose to use for migration of the database. It looks like I am able to both in GO and SQL files write a migration and then apply it via `goose up`.
Which could be used in a CI/CD flow which is my goal. 


## Goals
My goal is to in a CI/CD flow deploy a temporary database, apply migrations and seed some data. If posible, make a copy of the existing database and run the db and apply migrations and then run tests on the new migrations to look if everything looks good.


