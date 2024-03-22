build:
	go build -o ./bin/main ./cmd/*.go

run: tailwindcss templ build


deploy:
	git add .
	git commit -m "${opt:commit}"
	git push origin master

tailwindcss:
	npx tailwindcss -i tailwind-input.css -o static/css/tailwind.css


templ:
	~/go/bin/templ generate