CREATE TABLE if not exists challenge_participation (
                                          id uuid default gen_random_uuid() PRIMARY KEY,
                                          challenge_id UUID not null ,
                                          user_id UUID,
                                          status VARCHAR(20) NOT NULL,
                                          recycled_items_count INTEGER DEFAULT 0,
                                          joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                          deleted_at bigint default 0
);
