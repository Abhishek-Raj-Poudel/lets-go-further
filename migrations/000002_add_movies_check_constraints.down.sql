ALTER TABLE movies DROP CONSTRAINT IF EXISTS movies_runtime_check;

ALTER TABLE movies DROP CONSTRAINT IF EXISTS movies_year_check;

Alter TABLE movies DROP CONSTRAINT IF EXISTS genres_length_check;
