CREATE TYPE role as ENUM('super_admin','admin','user');

CREATE TABLE IF NOT EXISTS register(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phonenummer VARCHAR(50) NOT NULL,
    password TEXT DEFAULT '',
    image TEXT DEFAULT 'https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png',
    role role NOT NULL DEFAULT 'user',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS adminn(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES register(id),
    role role not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
)