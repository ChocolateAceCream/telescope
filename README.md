# telescope

## Frontend
react + tailwind + typescript + vite

steps:
1. install latest nodejs
2. npm install
3. change env variables in .env.development
4. npm run dev

## Backend
sqlc + pgsql + gin

steps:
1. install latest go
2. install sqlc
3. go mod tidy
4. change env variables in config.yaml
5. go run main.go


# Dev Journal
## 2025/01/28
Happy Chinese New Year!
- [x] add skeleton backend
- [x] add db
- [x] add docker for kafka

## 2025/01/29
- [x] add translations in backend, so all response can be translated based on user's language setting

## 2025/01/30
- [x] add translations in frontend
- [x] integrate zustand as state store
- [x] add frontend route guard

## 2025/01/31
- [x] config tailwind
tailwind 4 separated postcss to its own package, and old setup is not working anymore. After hours of struggle with v4, finally I switch back to tailwind v3 and everything is back to normal.

## 2025/02/01
- [x] add 404 page
- [x] add header nav bar