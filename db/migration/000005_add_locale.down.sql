-- Drop the trigger first
DROP TRIGGER IF EXISTS locale_updated_at ON public.locale;
-- Drop the table
DROP TABLE IF EXISTS public.locale;