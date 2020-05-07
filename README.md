# Fizz-Buzz

## Installation

Installer les dépendences :

```
go get -u github.com/go-chi/chi
go get github.com/stretchr/testify
```

Pensez à vérifier la valeur de votre variable d'environnement `GOPATH`.

Pour tester l'application :

```
go test ./... -v
```

## Lancement

Lancer l'application' :

```
go build && ./fizz-buzz
```

## Appel à l'API

Rendez vous à l'adresse suivante :

```
http://localhost:8080/fizz-buzz?int1=2&int2=3&limit=10&str1=fizz&str2=buzz
```