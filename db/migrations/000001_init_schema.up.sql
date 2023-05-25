CREATE TABLE "product" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(100),
  "price" float,
  "description" text,
  "image" text
);

CREATE TABLE "order" (
  "id" SERIAL PRIMARY KEY,
  "date" timestamp DEFAULT (now()),
  "customer_id" integer,
  "status" varchar(50)
);

CREATE TABLE "order_detail" (
  "id" SERIAL PRIMARY KEY,
  "order_id" integer,
  "product_id" integer
);

CREATE TABLE "customer" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(100),
  "email" varchar(100),
  "password" varchar(100),
  "created_at" timestamp DEFAULT (now())
);

CREATE INDEX "status_idx" ON "order" ("status");

CREATE INDEX "product_name_idx" ON "product" ("name");

CREATE INDEX "customer_idx" ON "order" ("customer_id");

CREATE UNIQUE INDEX ON "customer" ("email");

ALTER TABLE "order" ADD FOREIGN KEY ("customer_id") REFERENCES "customer" ("id");

ALTER TABLE "order_detail" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE "order_detail" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");
