CREATE TABLE if not exists eco_challenges (
                                id uuid default gen_random_uuid() PRIMARY KEY,
                                title VARCHAR(100) NOT NULL,
                                description TEXT,
                                start_date TIMESTAMP WITH TIME ZONE NOT NULL,
                                end_date TIMESTAMP WITH TIME ZONE NOT NULL,
                                reward_points INTEGER NOT NULL,
                                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                deleted_at bigint default 0
);
