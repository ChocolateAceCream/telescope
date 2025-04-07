-- Drop the trigger first
DROP TRIGGER IF EXISTS sketch_updated_at ON public.sketch;
-- Drop the table
DROP TABLE IF EXISTS public.sketch;