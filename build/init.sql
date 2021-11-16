CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "balance" NUMERIC NOT NULL,
  "currency" varchar NOT NULL DEFAULT 'RUB',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "deposits" (
  "id" bigserial PRIMARY KEY,
  "from_user_id" varchar DEFAULT 'atm machine',
  "to_user_id" varchar NOT NULL,
  "amount" NUMERIC NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "comment" varchar DEFAULT 'deposit'
);

CREATE TABLE "withdrawals" (
  "id" bigserial PRIMARY KEY,
  "from_user_id" varchar NOT NULL,
  "to_user_id" varchar DEFAULT 'atm machine',
  "amount" NUMERIC NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "comment" varchar DEFAULT 'withdrawal'
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "from_user_id" varchar NOT NULL,
  "to_user_id" varchar NOT NULL,
  "amount" NUMERIC NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "comment" varchar DEFAULT 'money transfer'
);

ALTER TABLE "deposits" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");

ALTER TABLE "withdrawals" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");