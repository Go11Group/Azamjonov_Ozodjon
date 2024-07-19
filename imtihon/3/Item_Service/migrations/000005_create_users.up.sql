CREATE TABLE if not exists recycling_centers (
                                   id uuid default gen_random_uuid() PRIMARY KEY,
                                   name VARCHAR(100) NOT NULL,
                                   address TEXT NOT NULL,
                                   accepted_materials JSONB NOT NULL,
                                   working_hours TEXT,
                                   contact_number VARCHAR(20),
                                   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                   deleted_at bigint default 0
);
