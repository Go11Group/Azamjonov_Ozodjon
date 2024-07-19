CREATE TABLE if not exists swaps (
                       id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                       offered_item_id UUID not null ,
                       requested_item_id UUID not null ,
                       requester_id UUID,
                       category_id uuid,
                       owner_id UUID,
                       status VARCHAR(20) NOT NULL,
                       message TEXT,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       completed_at TIMESTAMP WITH TIME ZONE,
                       deleted_at BIGINT DEFAULT 0
);
