# Secrets da pipeline

Os workflows de pull request compilam, testam, lintam e escaneiam sem estes secrets. O push da `main` só publica a imagem e atualiza o GitOps se os três existirem em **Settings → Secrets and variables → Actions** do repositório `souza-wallace/togglemaster`.

| Secret | Uso |
|---|---|
| `AWS_ACCESS_KEY_ID` | Login no ECR da conta `980228515477`, região `us-east-1` |
| `AWS_SECRET_ACCESS_KEY` | Par da chave acima |
| `GITOPS_TOKEN` | PAT com permissão de escrita em `souza-wallace/togglemaster-k8s` |

Os repositórios ECR `auth-service`, `flag-service`, `targeting-service`, `evaluation-service` e `analytics-service` precisam existir antes do primeiro push. A tag gravada no `deployment.yaml` do GitOps tem o formato `v1.0.0-<7 caracteres do SHA>`.
