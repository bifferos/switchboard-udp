.PHONY: build run clean run install uninstall test


APP_NAME := switchboard-udp

BINARY_NAME := $(APP_NAME)
INSTALL_DIR := /usr/local/bin
SERVICE_FILE := $(APP_NAME).service
SYSTEMD_DIR := /etc/systemd/system

USR_SHARE := /usr/share/$(APP_NAME)

STATE_DIR := /var/lib/switchboard


run: build
	./$(BINARY_NAME) -config ./config.json


install: build
	sudo install -m 755 $(BINARY_NAME) $(INSTALL_DIR)/
	sudo install -d $(STATE_DIR)
	sudo install -m 644 $(SERVICE_FILE) $(SYSTEMD_DIR)/
	sudo systemctl daemon-reexec
	sudo systemctl enable --now $(BINARY_NAME).service
	@echo "Installed and started $(BINARY_NAME).service"


uninstall:
	sudo systemctl disable --now $(APP_NAME).service
	sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	sudo rm -f $(SYSTEMD_DIR)/$(SERVICE_FILE)
	sudo systemctl daemon-reexec
	@echo "Uninstalled $(BINARY_NAME)"


info:
	systemctl status $(SERVICE_FILE)
	
	
logs:
	journalctl -u $(SERVICE_FILE) -n 20
	
	
stop:
	sudo systemctl stop $(SERVICE_FILE)


start:
	sudo systemctl start $(SERVICE_FILE)
	
	
tcpdump:
	sudo tcpdump -i any udp port 6060
	

build:
	go build -o $(BINARY_NAME) main.go


clean:
	rm -f $(BINARY_NAME) *~


test:
	./test.py


test-running:
	sudo ./test-running.py
	