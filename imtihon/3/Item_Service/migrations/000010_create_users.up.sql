CREATE TABLE if not exists eco_tips (
                          id uuid default gen_random_uuid() PRIMARY KEY,
                          title VARCHAR(100) NOT NULL,
                          content TEXT NOT NULL,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          deleted_at bigint default 0
);
