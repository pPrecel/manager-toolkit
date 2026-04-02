# create-k3d-cluster

A GitHub action that creates a local k3d Kubernetes cluster suitable for integration testing.

## How It Works

1. Installs k3d.
2. Optionally writes a `registries.yaml` config that mirrors `europe-docker.pkg.dev` through an authenticated proxy (only when `restricted-registry-sa` is provided).
3. Creates a cluster named `kyma` (kubeconfig context: `k3d-kyma`) with:
   - 1 agent node running k3s.
   - Ports 80 and 443 forwarded through the load balancer.
   - Flannel CNI disabled in favour of Calico (see below).
   - Traefik ingress controller disabled.
4. Creates the `kyma-system` namespace.
5. Installs Calico as the CNI (CRDs → Tigera operator → custom resources), waits for CoreDNS to become ready, and patches the Calico installation to use the same CNI binary and config paths as the default k3s Flannel setup.

## Inputs

| Input                    | Required | Description                                                                                                                                                                          |
| ------------------------ | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `restricted-registry-sa` | No       | Base64-encoded JSON key for a GCP Service Account with pull access to `europe-docker.pkg.dev`. When supplied, the cluster is configured to authenticate against restricted registry. |

## Usage

### Without a restricted registry

```yaml
- uses: actions/checkout@v4

- uses: kyma-project/manager-toolkit/.github/actions/create-k3d-cluster@main
```

### With a restricted registry

```yaml
- uses: actions/checkout@v4

- uses: kyma-project/manager-toolkit/.github/actions/create-k3d-cluster@main
  with:
    restricted-registry-sa: ${{ secrets.DOCKER_REGISTRY_SA_KEY }}
```

The `restricted-registry-sa` value must be a base64-encoded JSON service account key (i.e. the output of `base64 -w0 key.json`).

## Outputs

This action has no outputs. After it completes, `kubectl` and `k3d` commands will target the `k3d-kyma` context automatically.
