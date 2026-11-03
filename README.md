# ToggleMaster

Plataforma de feature flags. O código dos cinco microsserviços vive neste monorepo, mas o build e o deploy de cada um são independentes.

| Serviço | Linguagem | Porta | Papel |
|---|---|---|---|
| `auth-service` | Go | 8001 | Cria e valida chaves de API |
| `flag-service` | Python | 8002 | CRUD das definições de flags |
| `targeting-service` | Python | 8003 | Regras de segmentação |
| `evaluation-service` | Go | 8004 | Avalia a flag para o cliente e publica o evento |
| `analytics-service` | Python | 8005 | Consome os eventos e grava no DynamoDB |

Para subir tudo localmente: `docker compose up --build`.

## Deploy separado por path

Cada serviço tem o próprio workflow em `.github/workflows/` (`auth.yml`, `flag.yml`, `targeting.yml`, `evaluation.yml`, `analytics.yml`). Os workflows reutilizam `ci-go.yml` ou `ci-python.yml`.

O gatilho é um push ou um pull request na branch `main`, e só quando o commit altera o path daquele serviço. Um commit que mexe só em `auth-service/` não constrói nem publica os outros quatro. O mesmo vale para `flag-service/`, `targeting-service/`, `evaluation-service/` e `analytics-service/`.

A publicação da imagem no ECR e a atualização da tag no repositório GitOps (`souza-wallace/togglemaster-k8s`) acontecem apenas no push da `main`. Pull request roda o pipeline, sem publicar. O Argo CD, nesse outro repositório, sincroniza o cluster a partir da tag gravada.
