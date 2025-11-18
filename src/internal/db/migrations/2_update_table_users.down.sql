ALTER TABLE users
DROP CONSTRAINT users_profile_picture_fkey;

ALTER TABLE users
DROP COLUMN profile_picture;
