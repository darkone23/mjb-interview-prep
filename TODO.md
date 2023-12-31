TODO:

- [ x ] Datastore was checked into project
  - We do not want our production data to be stored in our git repository
    - Instead each unique deploy will require its own configuration
- [ x ] we don't need postgres or docker...
  - external data dependencies bring extra custodial complexity
    - is the DB local to us? or do we need to establish a secure connection?
  - let's just use sqlite3: fast enough until we scale to the megabucks
    - simpler to run today, easy to swap out later
    - considerations: will be harder if we have an existing dataset
      - would require a careful migration path
- [ x ] AddUser impl is wrong
  - using Query instead of Statement for insert
    - cannot query a null id...
- [ x ] SQL injection vuln in internal/user/service.go
  - Need to use prepared statements
- [ x ] better http routing
  - using gin so we don't have to roll this all ourselves
- [ x ] should probably use an ORM
  - enforce struct aligns with database
    - let's use https://github.com/volatiletech/sqlboiler
- [ x ] usernames should be unique
  - current migration has no such requirement
- [ x ] proper app configuration
  - where is the database located?
    - can probably use sqlboiler stuff for this
- [ x ] application should maintain a database connection
  - today each call to svc.AddUser opens a new database connection
    - instead we should have a shared connection pool
    - see https://turriate.com/articles/making-sqlite-faster-in-go for sqlite3 details
      - as written today will throw lock contention on sqlite3
- [ x ] better error handling
  - need at least some logging stuff so we can debug our errors
    - better would be proper instrumentation / observability
- [ ] implement other basic REST operations for User
  - [ x ] GET /user 
  - [ x ] GET /user/:id
  - [ ] DELETE /user/:id
  - [ ] PUT /user/:id
- [ ] basic test coverage
  - set up build action to test on commit
- [ ] implement signup and authentication workflow
  - use JWT auth
- [ ] do not store or send plaintext passwords over HTTP
  - vulnerable to mitm attacks between client and server
  - storing plaintext makes us vulnerable to database compromises
  - solution: clients send us an argon2 hash to store
    - it is the clients problem how they get it
      - libargon2, argon2-browser are good sources
    - server will not store this client generated key but will store a hash of it
      - auth requirements: 
        - client creates and sends a consistent hash
        - server creates a consistent hash from this and checks against what is stored in DB
- [ x ] wire up frontend
- [  ] basic html+js frontend
  - [ ] signup
  - [ ] login
  - [ ] listing

General questions:

- What is the context and use case of our service?
  - Do we need authentication?
    - Is it assumed anyone who can connect to our socket is already authenticated?
  - Do we need access controls?
    - Are there some operations only certain users can do?
- What is our desired user interface?
  - We are already over HTTP - are we a remote web API?
  - Do we want an HTML interface?
  - Do we want a command line text interface?
