# cda

Continuous Deployment Agent for Kubernetes

## Configuration

### Agent

The Agent must be configured by environment variables.

Environment Variable|Description                                                    |Type
--------------------|---------------------------------------------------------------|--------
`GITHUB_SECRET`     |The secret for a <br> GitHub Webhook.                          |*String*

## Install the service to Kubernetes with Helm

```bash
helm upgrade \
  --install cda deployments/chart/ \
  --set env.githubSecret=<your-github-secret> \
  --namespace cda \
  --force
```
