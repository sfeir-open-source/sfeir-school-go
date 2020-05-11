# Les dépendances Go

## Un workflow

- **go mod init** crée un nouveau module, initialiser le go.mod
- **go build, go test** comme d’habitude
- **go list -m all** affiche les dépendances du module courant
- **go get** change ou ajoute une dépendance
- **go mod** tidy nettoie les dépendances obsolètes
