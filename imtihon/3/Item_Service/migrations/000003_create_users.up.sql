CREATE TABLE if not exists items (
                       id uuid default gen_random_uuid() PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       description TEXT,
                       category_id UUID not null ,
                       condition VARCHAR(20) NOT NULL,
                       swap_preference JSONB,
                       owner_id UUID ,
                       status VARCHAR(20) NOT NULL,
                       listed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       deleted_at bigint default 0
);
