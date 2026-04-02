# configure-kubeconfig-clusterrole

A GitHub action that creates a dedicated kubeconfig context scoped to a given ClusterRole and switches to it.

## How It Works

1. Creates a ServiceAccount named after the requested ClusterRole in the `kube-system` namespace.
2. Binds that ServiceAccount to the ClusterRole via a ClusterRoleBinding.
3. Mints a short-lived token (24 h) for the ServiceAccount.
4. Registers the token as a new kubeconfig credential and context, then switches the active context to it.

All RBAC operations are performed against the admin context (`k3d-kyma`), which is always set by the [`create-k3d-cluster`](../create-k3d-cluster) action. Subsequent steps in the pipeline therefore run under the permissions of the specified ClusterRole.

## Prerequisites

| Requirement     | Details                                                                                                                                                                       |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **k3d cluster** | A running cluster whose admin context is named `k3d-kyma`. Use [`kyma-project/manager-toolkit/.github/actions/create-k3d-cluster`](../create-k3d-cluster) before this action. |
| **kubectl**     | Must be available in `PATH`.                                                                                                                                                  |

## Inputs

| Input         | Required | Description                                                                                                                                                 |
| ------------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `clusterrole` | Yes      | Name of the ClusterRole to bind to the pipeline ServiceAccount. Also used as the name of the new kubeconfig context (e.g. `edit`, `view`, `cluster-admin`). |

## Usage

```yaml
- uses: actions/checkout@v4

- uses: kyma-project/manager-toolkit/.github/actions/create-k3d-cluster@main

- uses: kyma-project/manager-toolkit/.github/actions/configure-kubeconfig-clusterrole@main
  with:
    clusterrole: view

- name: Steps here run as the 'view' ClusterRole
  run: kubectl get pods -A
```

Switching between roles within the same job is supported — call the action again with a different `clusterrole` value and it will create a new context and switch to it.
