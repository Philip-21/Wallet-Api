start_backend:
	cd backend && go run ./cmd/web

start_frontend:
	cd frontend && npm start

start_app:
	cd backend && go run ./cmd/web & \
	cd frontend && npm start
