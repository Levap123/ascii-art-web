# Ascii Art Web

Web application for creating an ASCII pattern


## Authors

- [@Samat](https://01.alem.school/git/ssadvaka)
- [@Pavel](https://01.alem.school/git/ssadvaka)

## Usage and Installation
Clone git repository and go to program directory:

```bash
git clone git@git.01.alem.school:LevapMik/ascii-art-web.git
cd ascii-art-web
```
Different ways how to run our web app:

 ```bash
go run cmd/main.go
```
```bash
make go
```
or via docker
```bash
make build && make run
```
then go to: http://localhost:8080

## Implementation details

Go programming language was used to create a web application. External packages were not used. Main package - net/http, for frontend we are execute static html pages.
