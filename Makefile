build:
	go build -o ./bin/main ./cmd/cbt-app-v1/main.go

run: tailwindcss templ build


deploy:
	git add .
	git commit -m "${opt:commit}"
	git push origin master

tailwindcss:
	bun run tailwindcss --config tailwind.config.js -i tailwind-input.css -o static/css/tailwind.css


templ:
	~/go/bin/templ generate