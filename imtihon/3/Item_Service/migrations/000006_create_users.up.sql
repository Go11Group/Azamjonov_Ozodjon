CREATE TABLE if not exists recycling_submissions (
                                       id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
                                       center_id UUID NOT NULL ,
                                       user_id UUID,
                                       eco_points_earned INTEGER NOT NULL,
                                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                       deleted_at BIGINT DEFAULT 0
);

CREATE TABLE if not exists recycling_submission_items (
                                            id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
                                            submission_id UUID NOT NULL ,
                                            item_id UUID,
                                            weight FLOAT NOT NULL,
                                            material TEXT NOT NULL
);
