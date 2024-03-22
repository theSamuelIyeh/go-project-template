CREATE TABLE public.profiles (
    id UUID PRIMARY KEY,
    user_id UUID references auth.users.id,      
)