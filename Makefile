ifneq (,$(wildcard ./goapp.env))
    include goapp.env
    export
endif

ifneq (,$(wildcard ./postgres.env))
    include postgres.env
    export
endif

export TAG=development

up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose up --build -d
