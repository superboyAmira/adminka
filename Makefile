USERNAME := $(shell whoami)
APP_NAME := adminka
APP_PORT := 5000
REPO_URL := https://github.com/superboyAmira/adminka
APP_DIR := /home/$(USERNAME)/app
DEPLOY_DIR := /var/www/app
REPO_DIR := /home/$(USERNAME)/adminka

# sed -i 's/^    /\t/' Makefile

update:
	sudo apt update && sudo apt upgrade -y

go_install: update
	wget https://go.dev/dl/go1.23.2.linux-amd64.tar.gz
	sudo rm -rf /usr/local/go
	sudo tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> ~/.profile
	export PATH=$$PATH:/usr/local/go/bin
	go version || true

git_install: go_install
	sudo apt install git -y

make_deploy_dirs: git_install
	mkdir -p $(APP_DIR)
	cd $(APP_DIR)
	sudo mkdir -p $(DEPLOY_DIR)
	sudo chown $(USERNAME):$(USERNAME) $(DEPLOY_DIR)
	sudo chmod 755 $(DEPLOY_DIR)

app_build: make_deploy_dirs
	cd $(REPO_DIR)/cmd/$(APP_NAME)/ && go build -o $(APP_NAME)
	mv $(REPO_DIR)/cmd/$(APP_NAME)/$(APP_NAME) $(DEPLOY_DIR)/

service_build: app_build
	service_build: app_build
	@sudo mv adminka.service /etc/systemd/system/$(APP_NAME).service
service_start: service_build
	sudo systemctl daemon-reload
	sudo systemctl enable $(APP_NAME).service
	sudo systemctl start $(APP_NAME).service
	sudo systemctl status $(APP_NAME).service

nginx_create: service_start
	sudo apt install nginx -y
	@sudo mv adminka /etc/nginx/sites-available/$(APP_NAME)
	sudo ln -sf /etc/nginx/sites-available/$(APP_NAME) /etc/nginx/sites-enabled/$(APP_NAME)
	sudo rm -f /etc/nginx/sites-enabled/default
	sudo nginx -t
	sudo systemctl restart nginx

test:
	@echo "=== TEST: HealthCheck ==="
	curl -s http://localhost/health | jq
	@echo "=== TEST: Create Player ==="
	curl -s -X POST http://localhost/players/create -H "Content-Type: application/json" -d '{"name":"Magnus","rating":2800}' | jq
	@echo "=== TEST: List Players ==="
	curl -s http://localhost/players | jq

all: nginx_create
