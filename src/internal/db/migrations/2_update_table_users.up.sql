ALTER TABLE users 
ADD COLUMN profile_picture INT 
REFERENCES media (id);

