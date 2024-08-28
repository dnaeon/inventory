CREATE TABLE IF NOT EXISTS "gcp_instance" (
    "name" varchar NOT NULL,
    "hostname" varchar NOT NULL,
    "instance_id" bigint NOT NULL,
    "project_id" varchar NOT NULL,
    "zone" varchar NOT NULL,
    "region" varchar NOT NULL,
    "can_ip_forward" boolean NOT NULL,
    "cpu_platform" varchar NOT NULL,
    "creation_timestamp" varchar NOT NULL,
    "description" varchar NOT NULL,
    "last_start_timestamp" varchar,
    "last_stop_timestamp" varchar,
    "last_suspend_timestamp" varchar,
    "machine_type" varchar NOT NULL,
    "min_cpu_platform" varchar NOT NULL,
    "self_link" varchar NOT NULL,
    "source_machine_image" varchar NOT NULL,
    "status" varchar NOT NULL,
    "status_message" varchar NOT NULL,
    "id" bigserial NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    CONSTRAINT "gcp_instance_key" UNIQUE ("instance_id", "project_id")
);
