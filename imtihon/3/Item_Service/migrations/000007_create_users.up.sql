CREATE TABLE if not exists ratings (
                         id uuid default gen_random_uuid() PRIMARY KEY,
                         user_id UUID,
                         rater_id UUID,
                         rating DECIMAL(2, 1) NOT NULL,
                         comment TEXT,
                         swap_id UUID not null ,
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         deleted_at bigint default 0
);
