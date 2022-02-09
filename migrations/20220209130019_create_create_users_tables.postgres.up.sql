CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
	name CHARACTER(70) NOT NULL,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	gender TEXT NOT NULL CHECK (gender IN ('MALE', 'FEMALE')),
  	otp TEXT,
	email_verified BOOLEAN DEFAULT false
  );