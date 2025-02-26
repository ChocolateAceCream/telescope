-- 000004_add_a_users_table.down

Drop Table If Exists "public"."password_logins";
Drop trigger If Exists a_users_updated_at_autocomplete On "public"."password_logins";