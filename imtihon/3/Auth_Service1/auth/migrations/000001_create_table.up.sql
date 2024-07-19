CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,

    bio TEXT,
    eco_points int default 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at bigint DEFAULT 0 
);
CREATE TABLE ecopoints (
                           id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                           user_id UUID NOT NULL,
                           points INTEGER NOT NULL,
                           reason TEXT NOT NULL,
                           timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                           type TEXT NOT NULL
);
CREATE TABLE token_blacklist (
                                 id SERIAL PRIMARY KEY,
                                 token TEXT NOT NULL,
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);