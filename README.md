## user-balance / Avito autumn 2021 backend trainee test

Projects starts with a command: make dev

API endpoints:

- /balance?user=userID 

(example http://localhost:4000/balance?user=aaaa)
- /balance?user=userID&currency=XXX

(http://localhost:4000/balance?user=aaaa&currency=USD)
- /deposit?user=userID&amount=xxxx

(curl -X POST http://localhost:4000/deposit?user=aaaa&amount=1000)
- /withdraw

(curl -X POST http://localhost:4000/withdraw?user=cccc&amount=1000)
- /transfer?from_user=userID&to_user=userID&amount=xxxx

(curl -X POST http://localhost:4000/transfer?from_user=cccc&to_user=bbbb&amount=xxxx)

General assumptions:
Since excangerateapi for free-tier accounts:
- provides only historic data -> we are making the calls using the "yesterday" date
- allows to use only EUR as a base currency -> calls are made for EUR to RUB and EUR to "requested currency", after that excange rate between RUB and "currency" are calculated
