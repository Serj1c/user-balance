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
