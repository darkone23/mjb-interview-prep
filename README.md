# golang interview prep

## Goal of this repo.
 
This repo is a fork of https://github.com/MatthewJamesBoyle/golang-interview-prep

The original repo was created as a coding exercise: It was partially generated 
by ChatGPT and did not follow best practices, contained multiple bugs and security issues.

The basics of this application: create a user and store it in the database

This repo is a modification of that repo to bring it up to date with web programming best practices, 
fix any security holes, and learn some go along the way!

## Changes made so far:

- Add HTTP routing library (gin)
  - fixes http response & marshalling issues
- Add ORM library (sqlboiler)
  - fixes sql vulnerabilities found in hand-rolled SQL
- Add build script (just)
- Update dev toolchain (devenv)
  - Add github CI
- Ported db from postgres to sqlite
  - ["one process programming"](https://crawshaw.io/blog/one-process-programming-notes)
- Use environment variables to connect to database
- Fix security issues in user table
  - Enforce unique usernames, unique passwords (!)
- Implement some basic REST methods
- Added web module for npm based frontend 

## Getting Started

Some dev tools are used to make life easier:

- [devenv](https://devenv.sh/)
- [direnv](https://direnv.net/)

Once both are installed you can get the application started by running `direnv allow && just setup-to-serve`

This will setup and migrate a local sqlite DB, code generate some schema files, and get your server running.

Once running the Go app, you can make a CURL request as follows:

```curl
 curl -X POST -H "Content-Type: application/json" -d '{"username":"john", "password":"secret"}' http://localhost:8080/user
```
