## user-balance / Avito autumn 2021 backend trainee test

Project starts with commands: 
1. ```make start``` (spins up docker containers)
2. ```make migrateup``` (creates all tables)

API endpoints:

1. Get user's balance: ```/balance?user=userID``` 
```
http://localhost:4000/balance?user=aaaa
```
2. Get user's balance in other currencies ```/balance?user=userID&currency=XXX```
```
http://localhost:4000/balance?user=aaaa&currency=USD
```
3. Deposit of money on a balance: ```/deposit?user=userID&amount=xxxx```
```
curl -X POST http://localhost:4000/deposit?user=aaaa&amount=1000
```
4. Withdrawal of money from a balance ```/withdraw?user=userID&amount=xxxx```
```
curl -X POST http://localhost:4000/withdraw?user=aaaa&amount=500
```
5. Transfer money from one user to another: ```/transfer?from_user=userID&to_user=userID&amount=xxxx```
```
curl -X POST http://localhost:4000/transfer?from_user=aaaa&to_user=bbbb&amount=100
```
6. List all operations with user's balance. Method accepts optional parameters: 

- for sorting: ```&sort=created_at&sort=asc```, 
```&sort=created_at&sort=desc```, ```&sort=amount&sort=asc```, ```&sort=amount&sort=desc```. 

```
http://localhost:4000/operations?user=userID&sort=amount&sort=desc
```
 
If left empty -> default will be ```&sort=created_at&sort=asc```

- for pagination: page - number of requested page, per_page - quantity of operations to be returned ```&page=xx&per_page=xx```

```
http://localhost:4000/operations?user=userID&page=1&per_page=5
```

If left empty -> default will be ```&page=1&per_page=10```

General assumptions:
1. Since excangerateapi for free-tier accounts:
- provides only historic data -> we are making the calls using the "yesterday" date
- allows using only EUR as a base currency -> two calls are made: the first one for EUR to RUB and the second one for EUR to "requested currency", after that an excange rate between RUB and "requested currency" are calculated and returned
2. First deposit of money on a balance creates a user. To receive money through a transaction a user must already be created
