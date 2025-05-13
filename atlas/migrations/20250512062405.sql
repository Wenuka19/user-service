-- Create "users" table
CREATE TABLE "public"."users" (
  "id" text NOT NULL,
  "name" text NULL,
  "email" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
