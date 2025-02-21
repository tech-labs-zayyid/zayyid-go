-- migrate:up
CREATE TABLE IF NOT EXISTS "testimoni" (
	"id" VARCHAR(50) NOT NULL,
	"user_name" VARCHAR(150) NOT NULL DEFAULT '',
	"position" VARCHAR(250) NOT NULL DEFAULT '',
	"deskripsi" TEXT NOT NULL,
	"photo_url" TEXT NOT NULL,
	"is_active" SMALLINT NOT NULL DEFAULT '0',
	"created_at" TIMESTAMP NULL DEFAULT NULL,
	"modified_at" TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY ("id")
);

CREATE INDEX index_testimoni ON testimoni (user_name);

-- migrate:down
DROP TABLE IF EXISTS testimoni CASCADE;