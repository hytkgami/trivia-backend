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

CREATE TABLE "public"."questions" (
    "question_id" varchar NOT NULL,
    "created_by" varchar NOT NULL, -- uid of user who created the question
    "lobby_id" varchar NOT NULL,
    "title" varchar NOT NULL,
    "order_number" int4 NOT NULL,
    "score" int4 NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY ("question_id"),
    FOREIGN KEY ("created_by") REFERENCES "public"."users"("uid") ON UPDATE CASCADE,
    FOREIGN KEY ("lobby_id") REFERENCES "public"."lobbies"("lobby_id") ON UPDATE CASCADE,
    UNIQUE ("lobby_id", "order_number"),
    UNIQUE ("lobby_id", "title")
);

CREATE TABLE "public"."answers" (
    "answer_id" varchar NOT NULL,
    "question_id" varchar NOT NULL,
    "uid" varchar NOT NULL,
    "content" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY ("answer_id"),
    FOREIGN KEY ("question_id") REFERENCES "public"."questions"("question_id") ON UPDATE CASCADE,
    FOREIGN KEY ("uid") REFERENCES "public"."users"("uid") ON UPDATE CASCADE
);
