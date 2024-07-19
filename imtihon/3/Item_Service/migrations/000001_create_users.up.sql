CREATE TABLE IF NOT EXISTS item_categories (
                                               id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                               name VARCHAR(50) NOT NULL,
                                               description TEXT,
                                               created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS recycled_items (
                                              id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                              category_id UUID NOT NULL,
                                              center_id UUID NOT NULL,
                                              user_id UUID NOT NULL,
                                              created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS eco_points (
                                          id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                          user_id UUID NOT NULL,
                                          eco_points INT NOT NULL,
                                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);