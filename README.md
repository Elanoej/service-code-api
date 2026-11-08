# service-code-api

API HTTP escrita em Go para consulta de códigos de serviços.

## Requisitos

- [Go](https://go.dev/dl/) 1.26 ou superior

## Como rodar

```bash
go run .
```

O servidor sobe em `http://localhost:8000`.

## Endpoints

### `GET /health`

Verifica se a API está no ar.

```bash
curl http://localhost:8000/health
```

```json
{ "status": "ok" }
```

### `GET /codigos?search=<termo>`

Lista os serviços cujo `codigo` ou `descricao` contenham o termo informado
(a busca é case-insensitive). Sem o parâmetro `search`, retorna todos.

```bash
curl "http://localhost:8000/codigos?search=programa"
```

```json
[
  {
    "codigo": "010201",
    "descricao": "Programação."
  }
]
```

## Testes

```bash
go test ./...
```

## Estrutura

```
.
├── main.go        # ponto de entrada, sobe o servidor na porta 8000
├── api.go         # rotas e handlers HTTP
├── api_test.go    # testes dos handlers
├── servico.go     # modelo Servico e dados mockados
└── utils/
    └── json.go    # helper para responder JSON
```
