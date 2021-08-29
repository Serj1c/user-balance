CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL DEFAULT 'RUB',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "deposits" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "withdrawls" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "from_user_id" varchar NOT NULL,
  "to_user_id" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "deposits" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "withdrawls" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");