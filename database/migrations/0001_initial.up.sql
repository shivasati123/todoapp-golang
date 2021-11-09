CREATE TABLE user_profile(
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             name TEXT NOT NULL,
                             email_id TEXT NOT NULL,
                             mobile_no TEXT ,
                             city TEXT NOT NULL,
                             is_admin BOOLEAN NOT NULL,
                             created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                             archived_at TIMESTAMP WITH TIME ZONE

);
