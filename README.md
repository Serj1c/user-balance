## user-balance / Avito autumn 2021 backend trainee test

Project starts with commands: 
1. make docker (spins up docker containers)
2. make migrateup (creates all tables)

API endpoints:

- Get user's balance: /balance?user=userID 

(example http://localhost:4000/balance?user=aaaa)

- Get user's balance in other currencies /balance?user=userID&currency=XXX

(http://localhost:4000/balance?user=aaaa&currency=USD)

- Deposit money on a balance: /deposit?user=userID&amount=xxxx

(curl -X POST http://localhost:4000/deposit?user=aaaa&amount=1000)

-  /withdraw?user=userID&amount=xxxx

(curl -X POST http://localhost:4000/withdraw?user=cccc&amount=1000)

- Transfer money from one user to another: /transfer?from_user=userID&to_user=userID&amount=xxxx

(curl -X POST http://localhost:4000/transfer?from_user=cccc&to_user=bbbb&amount=xxxx)

- List all operations with user's balance /operations

curl http://localhost:4000/operations?user=userID

General assumptions:
Since excangerateapi for free-tier accounts:
- provides only historic data -> we are making the calls using the "yesterday" date
- allows to use only EUR as a base currency -> calls are made for EUR to RUB and EUR to "requested currency", after that excange rate between RUB and "currency" are calculated
- 

<!-- TODO assumptions regarding floats with points -->

