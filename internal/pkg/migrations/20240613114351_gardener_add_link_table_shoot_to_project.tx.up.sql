CREATE TABLE IF NOT EXISTS "l_g_shoot_to_project" (
    "shoot_id" bigint NOT NULL,
    "project_id" bigint NOT NULL,
    "id" bigserial NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("shoot_id") REFERENCES "g_shoot" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("project_id") REFERENCES "g_project" ("id") ON DELETE CASCADE,
    CONSTRAINT "l_g_shoot_to_project_key" UNIQUE ("shoot_id", "project_id")
);
