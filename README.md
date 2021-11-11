## user-balance / Avito autumn 2021 backend trainee test

Project starts with commands: 
1. ```make start``` (spins up docker containers)
2. ```make migrateup``` (creates all tables)

API endpoints:

- Get user's balance: ```/balance?user=userID``` 
```
http://localhost:4000/balance?user=aaaa
```
- Get user's balance in other currencies ```/balance?user=userID&currency=XXX```
```
http://localhost:4000/balance?user=aaaa&currency=USD
```
- Deposit of money on a balance: ```/deposit?user=userID&amount=xxxx```
```
curl -X POST http://localhost:4000/deposit?user=aaaa&amount=1000
```
- Withdrawal of money from a balance ```/withdraw?user=userID&amount=xxxx```
```
curl -X POST http://localhost:4000/withdraw?user=aaaa&amount=500
```
- Transfer money from one user to another: ```/transfer?from_user=userID&to_user=userID&amount=xxxx```
```
curl -X POST http://localhost:4000/transfer?from_user=aaaa&to_user=bbbb&amount=100
```
- List all operations with user's balance. Method accepts optional parameters: &sort=amount&sort=desc (sorting by default on a date of operation in ascending order)
```
http://localhost:4000/operations?user=userID
```
- List all operations with user's balance
```
http://localhost:4000/operations?user=userID&sort=amount&sort=desc
```
General assumptions:
1. Since excangerateapi for free-tier accounts:
- provides only historic data -> we are making the calls using the "yesterday" date
- allows using only EUR as a base currency -> two calls are made: the first one for EUR to RUB and the second one for EUR to "requested currency", after that an excange rate between RUB and "requested currency" are calculated and returned
2. First deposit of money on a balance creates a user. To receive money through a transaction a user must already be created