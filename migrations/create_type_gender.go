package migrations

import "github.com/ElegantSoft/shabahy/db"

func CreateGenderType() {
	db.DB.Exec(`DO $$
						BEGIN
							IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender') THEN
								CREATE TYPE gender AS ENUM (
									'male', 'female'
								);
							END IF;
						END$$;`)
}
