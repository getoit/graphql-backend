# gqlgen with ent starter

This repository is a gqlgen with ent backend starter template.

## How to use

Let's assume your project is at github.com/user/repo

### 1. Clone the repository

```bash
mkdir -p ~/go/src/github.com/user
cd ~/go/src/github.com/user
git clone github.com/dlukt/graphql-backend-starter repo
cd repo
```

### 2. Replace the original repository name with your repository name

```bash
chmod +x ./update-schema.sh
./update-schema.sh github.com/user/repo
```

### 3. remove all "starter" occurences and replace with your repo name

```bash
rm graph/generated/starter.generated.go
mv starter.graphql repo.graphql
# remember, repo is your repo name
mv graph/starter.resolvers.go graph/repo.resolvers.go
```

edit gqlgen.yml and replace the `- starter.graphql` with your `- repo.graphql`

```yaml
schema:
  - ent.graphql
  - repo.graphql
```

Regenerate

```bash
go generate ./...
```

### 4. Change the git repo url

```bash
rm -rf .git
git init
git add .
git remote add origin github.com/user/repo
git commit -m 'initial'
git push -u origin master
```

## Adding new entities

```bash
alias ent='go run -mod=mod entgo.io/ent/cmd/ent'
```

`cd` into your project root.

```bash
ent new Entity # capitalization matters
```

add the new entity to `gqlgen.yml`

```yaml
autobind:
  - github.com/dlukt/graphql-backend-starter/ent
  - github.com/dlukt/graphql-backend-starter/ent/profile
  - github.com/dlukt/graphql-backend-starter/ent/entity
```

and edit the `ent/schema/entity.go` file.
Afterwards, regenerate all the things.

```bash
go generate ./...
```

## OIDC

### Claims

This package assumes Keycloak being the OIDC IDP.
Therefore the [claims object](rules/claims/claims.go) reflects Keycloak's claim structure.
Change this to your claim structure.
For instance I'm using Zitadel, adding per project grants.
The claims structure looks like this:

```go
type Claims struct {
    Aud   []string  `json:"aud"`
    Exp   time.Time `json:"exp"`
    Iat   time.Time `json:"iat"`
    Iss   string    `json:"iss"`
    Jti   string    `json:"jti"`
    Nbf   time.Time `json:"nbf"`
    Roles []string  `json:"roles"`
    Sub   string    `json:"sub"`
}
```

### Making read access require auth

On line 79 in `cmd/graphql.go` remove the `options.IsPermissive(),`.

## Further reading

[Ent.io GraphQL Tutorial](https://entgo.io/docs/tutorial-todo-gql)
