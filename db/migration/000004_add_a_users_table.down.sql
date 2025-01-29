-- 000004_add_a_users_table.down

Drop Table If Exists "public"."a_users";
Drop trigger If Exists a_users_updated_at_autocomplete On "public"."a_users";