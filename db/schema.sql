CREATE TABLE "public"."users" (
    "uid" varchar NOT NULL,
    "name" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY ("uid")
);
