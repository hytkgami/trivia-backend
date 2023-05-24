CREATE TABLE "public"."users" (
    "uid" varchar NOT NULL,
    "name" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY ("uid")
);

CREATE TABLE "public"."lobbies" (
    "lobby_id" varchar NOT NULL,
    "owner_uid" varchar NOT NULL,
    "name" varchar NOT NULL,
    "is_public" bool NOT NULL DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY ("lobby_id"),
    FOREIGN KEY ("owner_uid") REFERENCES "public"."users"("uid") ON UPDATE CASCADE
);
