-- Drop the trigger first
DROP TRIGGER IF EXISTS project_updated_at ON public.project;
-- Drop the table
DROP TABLE IF EXISTS public.project;