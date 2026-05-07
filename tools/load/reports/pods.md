# Pods

## Summary

Pods with memory growth exceeding 200 Mi from baseline to 5000 resources:

| pod (namespace/name)                                                        | baseline | 500 res | 1000 res | 2000 res | 3000 res | 4000 res | 5000 res | delta    |
| --------------------------------------------------------------------------- | -------- | ------- | -------- | -------- | -------- | -------- | -------- | -------- |
| kyma-system/rma-kube-state-metrics-6879cd7599-q9254                         | 1827 Mi  | 1762 Mi | 1952 Mi  | 2826 Mi  | 3206 Mi  | 4146 Mi  | 4179 Mi  | +2352 Mi |
| korifi/korifi-controllers-controller-manager-5ddd86668-85tk6                | 45 Mi    | 166 Mi  | 286 Mi   | 724 Mi   | 804 Mi   | 1074 Mi  | 2363 Mi  | +2318 Mi |
| kyma-system/api-gateway-controller-manager-66d9d4d67d-f867w                 | 73 Mi    | 195 Mi  | 370 Mi   | 574 Mi   | 873 Mi   | 1258 Mi  | 2070 Mi  | +1997 Mi |
| kyma-system/rma-victoria-metrics-agent-77ff8d586c-h5zzp                     | 209 Mi   | 388 Mi  | 604 Mi   | 1029 Mi  | 1366 Mi  | 1832 Mi  | 1750 Mi  | +1541 Mi |
| kyma-system/istio-controller-manager-6b7f78d558-fmlph                       | 63 Mi    | 184 Mi  | 387 Mi   | 633 Mi   | 666 Mi   | 1345 Mi  | 1345 Mi  | +1282 Mi |
| istio-system/istiod-66b4585c8b-njhj4                                        | 574 Mi   | 589 Mi  | 606 Mi   | 641 Mi   | 704 Mi   | 704 Mi   | 1678 Mi  | +1104 Mi |
| kyma-system/warden-operator-69bf7ff597-rv5bv                                | 43 Mi    | 157 Mi  | 314 Mi   | 574 Mi   | 685 Mi   | 1015 Mi  | 1015 Mi  | +972 Mi  |
| kpack/kpack-controller-5c6866cf46-twcx6                                     | 62 Mi    | 163 Mi  | 325 Mi   | 538 Mi   | 699 Mi   | 1005 Mi  | 1023 Mi  | +961 Mi  |
| kube-system/metrics-server-bd6684798-d57vf                                  | 532 Mi   | 538 Mi  | 598 Mi   | 604 Mi   | 764 Mi   | 843 Mi   | 946 Mi   | +414 Mi  |
| kube-system/metrics-server-bd6684798-psls2                                  | 406 Mi   | 420 Mi  | 463 Mi   | 427 Mi   | 413 Mi   | 542 Mi   | 626 Mi   | +220 Mi  |
| kyma-system/telemetry-metric-agent-d2rp9                                    | 111 Mi   | 122 Mi  | 128 Mi   | 163 Mi   | 178 Mi   | 211 Mi   | 344 Mi   | +233 Mi  |
| istio-system/istiod-66b4585c8b-rnqzx                                        | 447 Mi   | 447 Mi  | 458 Mi   | 486 Mi   | 504 Mi   | 558 Mi   | 648 Mi   | +201 Mi  |

Pods that have a valid baseline but start failing in further steps:

| pod (namespace/name)                                                        | baseline | 500 res | 1000 res | 2000 res | 3000 res | 4000 res | 5000 res |
| --------------------------------------------------------------------------- | -------- | ------- | -------- | -------- | -------- | -------- | -------- |
| cap-operator-system/cap-operator-controller-manager-57cdf94999-dspsd        | 44 Mi    | FAILING | FAILING  | -        | FAILING  | FAILING  | FAILING  |
| registry-proxy/registry-proxy-controller-5955cc9f9f-g82cd                   | 86 Mi    | 127 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |

## Full report

| pod (namespace/name)                                                        | baseline | 500 res | 1000 res | 2000 res | 3000 res | 4000 res | 5000 res |
| --------------------------------------------------------------------------- | -------- | ------- | -------- | -------- | -------- | -------- | -------- |
| antek-system/fake-metrics-746484997-jncvm                                   | 10 Mi    | 10 Mi   | 10 Mi    | 10 Mi    | 10 Mi    | 10 Mi    | 10 Mi    |
| antek-system/opensearch-external-scaler-fdbb7ffc-97zll                      | 8 Mi     | 8 Mi    | 8 Mi     | 8 Mi     | 8 Mi     | 8 Mi     | 8 Mi     |
| cap-operator-system/cap-operator-controller-9c9bc97bb-5kpsc                 | 22 Mi    | 22 Mi   | 22 Mi    | 22 Mi    | 22 Mi    | 22 Mi    | 26 Mi    |
| cap-operator-system/cap-operator-controller-manager-57cdf94999-dspsd        | 44 Mi    | FAILING | FAILING  | -        | FAILING  | FAILING  | FAILING  |
| cap-operator-system/cap-operator-subscription-server-d997f7d8f-frq6k        | 9 Mi     | 9 Mi    | 9 Mi     | 9 Mi     | 9 Mi     | 9 Mi     | 9 Mi     |
| cap-operator-system/cap-operator-webhook-786459c7ff-kl8pm                   | 11 Mi    | 11 Mi   | 11 Mi    | 11 Mi    | 11 Mi    | 11 Mi    | 11 Mi    |
| cfapi-system/btp-service-broker-cc965b778-qjjpz                             | 12 Mi    | 12 Mi   | 12 Mi    | 12 Mi    | 15 Mi    | 15 Mi    | 15 Mi    |
| cfapi-system/cfapi-operator-989bb945-4mn9s                                  | 135 Mi   | 135 Mi  | 135 Mi   | 136 Mi   | 135 Mi   | 136 Mi   | 136 Mi   |
| cfapi-system/contour-contour-57fd8d96b8-mn452                               | 36 Mi    | 37 Mi   | 38 Mi    | 38 Mi    | 38 Mi    | 38 Mi    | 39 Mi    |
| cfapi-system/contour-envoy-2v5vq                                            | 32 Mi    | 32 Mi   | 32 Mi    | 32 Mi    | 32 Mi    | 32 Mi    | 32 Mi    |
| cfapi-system/contour-envoy-7x65v                                            | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | 33 Mi    | 33 Mi    | 33 Mi    |
| cfapi-system/contour-envoy-xbjgh                                            | 34 Mi    | 34 Mi   | 34 Mi    | 34 Mi    | 34 Mi    | 34 Mi    | 34 Mi    |
| cfapi-system/contour-envoy-zdqns                                            | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | 35 Mi    | 35 Mi    | 35 Mi    |
| demo-app/http-echo-677d479d69-s24pf                                         | 78 Mi    | 79 Mi   | 79 Mi    | 79 Mi    | 75 Mi    | 78 Mi    | 78 Mi    |
| dev/api-postgresql-go-67f4797b5c-wppkk                                      | 69 Mi    | 69 Mi   | 73 Mi    | 73 Mi    | 70 Mi    | 70 Mi    | 74 Mi    |
| dev/commerce-mock-75648c645b-lh26b                                          | 189 Mi   | 189 Mi  | 184 Mi   | 184 Mi   | 184 Mi   | 184 Mi   | 186 Mi   |
| dev/fe-ui5-postgresql-55ff69fbd7-dhqgm                                      | 75 Mi    | 75 Mi   | 68 Mi    | 69 Mi    | 69 Mi    | 71 Mi    | 71 Mi    |
| docker-registry/dockerregistry-fd5f6685c-tfxws                              | 71 Mi    | 71 Mi   | 71 Mi    | 71 Mi    | 80 Mi    | 80 Mi    | 80 Mi    |
| docker-registry/dockerregistry-operator-587f69dbc4-rtz9z                    | 45 Mi    | 45 Mi   | 45 Mi    | 42 Mi    | 42 Mi    | 43 Mi    | 44 Mi    |
| istio-system/istio-cni-node-2fll2                                           | 23 Mi    | 24 Mi   | 24 Mi    | 23 Mi    | 23 Mi    | 33 Mi    | 36 Mi    |
| istio-system/istio-cni-node-gsf2b                                           | 37 Mi    | 37 Mi   | 37 Mi    | 37 Mi    | 37 Mi    | 37 Mi    | 36 Mi    |
| istio-system/istio-cni-node-h26x8                                           | 29 Mi    | 29 Mi   | 29 Mi    | 30 Mi    | 30 Mi    | 29 Mi    | 30 Mi    |
| istio-system/istio-cni-node-httjn                                           | 45 Mi    | 45 Mi   | 46 Mi    | 47 Mi    | 47 Mi    | 47 Mi    | 47 Mi    |
| istio-system/istio-cni-node-p6qxd                                           | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | 31 Mi    | 31 Mi    | 30 Mi    |
| istio-system/istio-cni-node-vzr4g                                           | 22 Mi    | 22 Mi   | 22 Mi    | 22 Mi    | 22 Mi    | 23 Mi    | 24 Mi    |
| istio-system/istio-cni-node-wsspg                                           | 30 Mi    | 30 Mi   | 31 Mi    | 31 Mi    | 32 Mi    | 32 Mi    | 32 Mi    |
| istio-system/istio-ingressgateway-76974d6f7f-65kbd                          | 98 Mi    | 98 Mi   | 95 Mi    | 100 Mi   | 95 Mi    | 94 Mi    | 100 Mi   |
| istio-system/istio-ingressgateway-76974d6f7f-c5h5v                          | 70 Mi    | 70 Mi   | 73 Mi    | 73 Mi    | 71 Mi    | 73 Mi    | 73 Mi    |
| istio-system/istio-ingressgateway-76974d6f7f-thkxr                          | 160 Mi   | 160 Mi  | 159 Mi   | 161 Mi   | 161 Mi   | 158 Mi   | 158 Mi   |
| istio-system/istiod-66b4585c8b-7ndn9                                        | -        | -       | -        | -        | -        | -        | 183 Mi   |
| istio-system/istiod-66b4585c8b-bdjhq                                        | -        | -       | -        | -        | 221 Mi   | -        | -        |
| istio-system/istiod-66b4585c8b-bdvvg                                        | -        | -       | -        | -        | -        | 220 Mi   | -        |
| istio-system/istiod-66b4585c8b-njhj4                                        | 574 Mi   | 589 Mi  | 606 Mi   | 641 Mi   | 704 Mi   | 704 Mi   | 1678 Mi  |
| istio-system/istiod-66b4585c8b-rnqzx                                        | 447 Mi   | 447 Mi  | 458 Mi   | 486 Mi   | 504 Mi   | 558 Mi   | 648 Mi   |
| istio-system/istiod-66b4585c8b-shhsj                                        | -        | -       | -        | -        | -        | -        | 180 Mi   |
| istio-system/istiod-66b4585c8b-tcpm9                                        | -        | -       | -        | -        | -        | -        | 181 Mi   |
| istio-system/istiod-66b4585c8b-xc8v2                                        | -        | -       | -        | -        | -        | 218 Mi   | -        |
| korifi/korifi-api-deployment-6d8c9f784d-lr6kh                               | 13 Mi    | 13 Mi   | 14 Mi    | 15 Mi    | 17 Mi    | 18 Mi    | 18 Mi    |
| korifi/korifi-controllers-controller-manager-5ddd86668-85tk6                | 45 Mi    | 166 Mi  | 286 Mi   | 724 Mi   | 804 Mi   | 1074 Mi  | 2363 Mi  |
| kpack/kpack-controller-5c6866cf46-twcx6                                     | 62 Mi    | 163 Mi  | 325 Mi   | 538 Mi   | 699 Mi   | 1005 Mi  | 1023 Mi  |
| kpack/kpack-webhook-54cf8f75bf-pk754                                        | 62 Mi    | 62 Mi   | 62 Mi    | 62 Mi    | 62 Mi    | 62 Mi    | 62 Mi    |
| kube-system/apiserver-proxy-479wm                                           | 31 Mi    | 31 Mi   | 31 Mi    | 31 Mi    | 31 Mi    | 31 Mi    | 31 Mi    |
| kube-system/apiserver-proxy-5lr8p                                           | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | 30 Mi    | 30 Mi    | 30 Mi    |
| kube-system/apiserver-proxy-8kw9m                                           | 25 Mi    | 25 Mi   | 25 Mi    | 25 Mi    | 25 Mi    | 25 Mi    | 25 Mi    |
| kube-system/apiserver-proxy-997dr                                           | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | 33 Mi    | 33 Mi    | 33 Mi    |
| kube-system/apiserver-proxy-hf6qf                                           | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | 30 Mi    | 30 Mi    | 30 Mi    |
| kube-system/apiserver-proxy-jpxjc                                           | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | 30 Mi    | 30 Mi    | 30 Mi    |
| kube-system/apiserver-proxy-q688t                                           | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | 33 Mi    | 33 Mi    | 33 Mi    |
| kube-system/blackbox-exporter-59694df6-tvrdm                                | 17 Mi    | 17 Mi   | 16 Mi    | 16 Mi    | 16 Mi    | 18 Mi    | 18 Mi    |
| kube-system/blackbox-exporter-59694df6-vwn7t                                | 14 Mi    | 15 Mi   | 16 Mi    | 17 Mi    | 17 Mi    | 14 Mi    | 15 Mi    |
| kube-system/calico-node-gw2x6                                               | 211 Mi   | 212 Mi  | 219 Mi   | 229 Mi   | 229 Mi   | 252 Mi   | 271 Mi   |
| kube-system/calico-node-jg47c                                               | 234 Mi   | 233 Mi  | 238 Mi   | 250 Mi   | 266 Mi   | 281 Mi   | 281 Mi   |
| kube-system/calico-node-kzvcd                                               | 198 Mi   | 200 Mi  | 204 Mi   | 211 Mi   | 230 Mi   | 253 Mi   | 240 Mi   |
| kube-system/calico-node-lmvlz                                               | 209 Mi   | 212 Mi  | 221 Mi   | 229 Mi   | 244 Mi   | 244 Mi   | 272 Mi   |
| kube-system/calico-node-ntsb5                                               | 202 Mi   | 198 Mi  | 203 Mi   | 211 Mi   | 227 Mi   | 242 Mi   | 264 Mi   |
| kube-system/calico-node-vertical-autoscaler-57d6f4f8f-t6p6x                 | 12 Mi    | 13 Mi   | 11 Mi    | 12 Mi    | 13 Mi    | 13 Mi    | 13 Mi    |
| kube-system/calico-node-vjp5x                                               | 126 Mi   | 132 Mi  | 145 Mi   | 158 Mi   | 164 Mi   | 190 Mi   | 212 Mi   |
| kube-system/calico-node-xfhnq                                               | 118 Mi   | 125 Mi  | 135 Mi   | 160 Mi   | 161 Mi   | 183 Mi   | 199 Mi   |
| kube-system/calico-typha-deploy-bd76c7b67-496gm                             | 111 Mi   | 115 Mi  | 124 Mi   | 130 Mi   | 137 Mi   | 144 Mi   | 154 Mi   |
| kube-system/calico-typha-deploy-bd76c7b67-5k8hj                             | 95 Mi    | 97 Mi   | 108 Mi   | 108 Mi   | 117 Mi   | 120 Mi   | 121 Mi   |
| kube-system/calico-typha-horizontal-autoscaler-66b75b5d95-rzspb             | 13 Mi    | 13 Mi   | 13 Mi    | 13 Mi    | 13 Mi    | 13 Mi    | 13 Mi    |
| kube-system/calico-typha-vertical-autoscaler-54b5769b9-8fjp2                | 12 Mi    | 12 Mi   | 12 Mi    | 13 Mi    | 13 Mi    | 13 Mi    | 13 Mi    |
| kube-system/cloud-node-manager-bw6nb                                        | 15 Mi    | 15 Mi   | 15 Mi    | 15 Mi    | 15 Mi    | 15 Mi    | 15 Mi    |
| kube-system/cloud-node-manager-d4qkw                                        | 17 Mi    | 17 Mi   | 17 Mi    | 17 Mi    | 17 Mi    | 17 Mi    | 17 Mi    |
| kube-system/cloud-node-manager-gmgxt                                        | 17 Mi    | 17 Mi   | 17 Mi    | 17 Mi    | 17 Mi    | 17 Mi    | 17 Mi    |
| kube-system/cloud-node-manager-k25vj                                        | 37 Mi    | 37 Mi   | 37 Mi    | 37 Mi    | 37 Mi    | 37 Mi    | 37 Mi    |
| kube-system/cloud-node-manager-krnf2                                        | 14 Mi    | 14 Mi   | 14 Mi    | 14 Mi    | 14 Mi    | 14 Mi    | 14 Mi    |
| kube-system/cloud-node-manager-lrzq5                                        | 28 Mi    | 28 Mi   | 28 Mi    | 28 Mi    | 29 Mi    | 29 Mi    | 29 Mi    |
| kube-system/cloud-node-manager-nzjvr                                        | 16 Mi    | 16 Mi   | 16 Mi    | 16 Mi    | 16 Mi    | 16 Mi    | 16 Mi    |
| kube-system/coredns-57dd88d9c5-4mr4k                                        | 35 Mi    | 35 Mi   | 35 Mi    | 35 Mi    | 36 Mi    | 36 Mi    | 30 Mi    |
| kube-system/coredns-57dd88d9c5-986ls                                        | 23 Mi    | 23 Mi   | 23 Mi    | 24 Mi    | 23 Mi    | 23 Mi    | 24 Mi    |
| kube-system/csi-driver-node-disk-62qkz                                      | 31 Mi    | 31 Mi   | 31 Mi    | 31 Mi    | 31 Mi    | 31 Mi    | 31 Mi    |
| kube-system/csi-driver-node-disk-7vzk5                                      | 38 Mi    | 38 Mi   | 38 Mi    | 38 Mi    | 38 Mi    | 38 Mi    | 38 Mi    |
| kube-system/csi-driver-node-disk-8vsm8                                      | 28 Mi    | 29 Mi   | 29 Mi    | 29 Mi    | 29 Mi    | 29 Mi    | 29 Mi    |
| kube-system/csi-driver-node-disk-gxksp                                      | 56 Mi    | 56 Mi   | 56 Mi    | 56 Mi    | 56 Mi    | 56 Mi    | 56 Mi    |
| kube-system/csi-driver-node-disk-hqrj7                                      | 34 Mi    | 34 Mi   | 34 Mi    | 34 Mi    | 34 Mi    | 34 Mi    | 34 Mi    |
| kube-system/csi-driver-node-disk-q7qns                                      | 31 Mi    | 31 Mi   | 31 Mi    | 31 Mi    | 31 Mi    | 31 Mi    | 31 Mi    |
| kube-system/csi-driver-node-disk-rpkkb                                      | 29 Mi    | 29 Mi   | 29 Mi    | 29 Mi    | 29 Mi    | 29 Mi    | 29 Mi    |
| kube-system/csi-driver-node-file-2zh57                                      | 27 Mi    | 27 Mi   | 27 Mi    | 27 Mi    | 27 Mi    | 27 Mi    | 27 Mi    |
| kube-system/csi-driver-node-file-49s9r                                      | 25 Mi    | 25 Mi   | 26 Mi    | 26 Mi    | 26 Mi    | 26 Mi    | 26 Mi    |
| kube-system/csi-driver-node-file-4cqsk                                      | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | 30 Mi    | 30 Mi    | 30 Mi    |
| kube-system/csi-driver-node-file-7tk8r                                      | 29 Mi    | 29 Mi   | 29 Mi    | 29 Mi    | 29 Mi    | 29 Mi    | 29 Mi    |
| kube-system/csi-driver-node-file-gfbhk                                      | 35 Mi    | 35 Mi   | 35 Mi    | 35 Mi    | 35 Mi    | 35 Mi    | 32 Mi    |
| kube-system/csi-driver-node-file-gx842                                      | 28 Mi    | 28 Mi   | 28 Mi    | 28 Mi    | 28 Mi    | 28 Mi    | 28 Mi    |
| kube-system/csi-driver-node-file-j5nqx                                      | 44 Mi    | 44 Mi   | 44 Mi    | 44 Mi    | 44 Mi    | 44 Mi    | 44 Mi    |
| kube-system/egress-filter-applier-bcc7b                                     | 14 Mi    | 14 Mi   | 13 Mi    | 13 Mi    | 13 Mi    | 13 Mi    | 13 Mi    |
| kube-system/egress-filter-applier-ddz6s                                     | 9 Mi     | 9 Mi    | 8 Mi     | 8 Mi     | 8 Mi     | 8 Mi     | 10 Mi    |
| kube-system/egress-filter-applier-fcwfv                                     | 11 Mi    | 11 Mi   | 9 Mi     | 9 Mi     | 9 Mi     | 9 Mi     | 9 Mi     |
| kube-system/egress-filter-applier-k8bxb                                     | 9 Mi     | 9 Mi    | 9 Mi     | 9 Mi     | 9 Mi     | 9 Mi     | 9 Mi     |
| kube-system/egress-filter-applier-mpcwc                                     | 16 Mi    | 16 Mi   | 13 Mi    | 13 Mi    | 13 Mi    | 13 Mi    | 13 Mi    |
| kube-system/egress-filter-applier-q5zr5                                     | 12 Mi    | 12 Mi   | 10 Mi    | 10 Mi    | 10 Mi    | 10 Mi    | 10 Mi    |
| kube-system/egress-filter-applier-t4g7n                                     | 13 Mi    | 13 Mi   | 12 Mi    | 12 Mi    | 12 Mi    | 12 Mi    | 12 Mi    |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-69l5f                           | 23 Mi    | 24 Mi   | 24 Mi    | 23 Mi    | 25 Mi    | 26 Mi    | 26 Mi    |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-cbd4s                           | 38 Mi    | 37 Mi   | 37 Mi    | 37 Mi    | 37 Mi    | 38 Mi    | 38 Mi    |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-hrsbg                           | 30 Mi    | 30 Mi   | 29 Mi    | 28 Mi    | 29 Mi    | 30 Mi    | 31 Mi    |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-lx7tt                           | 44 Mi    | 44 Mi   | 44 Mi    | 45 Mi    | 45 Mi    | 45 Mi    | 47 Mi    |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-n25mv                           | 30 Mi    | 30 Mi   | 27 Mi    | 29 Mi    | 30 Mi    | 30 Mi    | 31 Mi    |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-nkmzf                           | 29 Mi    | 29 Mi   | 29 Mi    | 28 Mi    | 28 Mi    | 29 Mi    | 31 Mi    |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-qlhfv                           | 35 Mi    | 35 Mi   | 34 Mi    | 35 Mi    | 35 Mi    | 36 Mi    | 34 Mi    |
| kube-system/metrics-server-bd6684798-d57vf                                  | 532 Mi   | 538 Mi  | 598 Mi   | 604 Mi   | 764 Mi   | 843 Mi   | 946 Mi   |
| kube-system/metrics-server-bd6684798-psls2                                  | 406 Mi   | 420 Mi  | 463 Mi   | 427 Mi   | 413 Mi   | 542 Mi   | 626 Mi   |
| kube-system/network-problem-detector-host-58tnx                             | 22 Mi    | 22 Mi   | 22 Mi    | 23 Mi    | 23 Mi    | 22 Mi    | 23 Mi    |
| kube-system/network-problem-detector-host-bz8jp                             | 21 Mi    | 21 Mi   | 21 Mi    | 22 Mi    | 22 Mi    | 23 Mi    | 22 Mi    |
| kube-system/network-problem-detector-host-gbwfs                             | 26 Mi    | 26 Mi   | 27 Mi    | 28 Mi    | 28 Mi    | 29 Mi    | 29 Mi    |
| kube-system/network-problem-detector-host-l6t7s                             | 23 Mi    | 23 Mi   | 23 Mi    | 23 Mi    | 22 Mi    | 23 Mi    | 23 Mi    |
| kube-system/network-problem-detector-host-p8728                             | 24 Mi    | 25 Mi   | 25 Mi    | 25 Mi    | 24 Mi    | 24 Mi    | 25 Mi    |
| kube-system/network-problem-detector-host-ts5jd                             | 16 Mi    | 16 Mi   | 16 Mi    | 17 Mi    | 18 Mi    | 19 Mi    | 19 Mi    |
| kube-system/network-problem-detector-host-xx48m                             | 23 Mi    | 23 Mi   | 22 Mi    | 23 Mi    | 23 Mi    | 24 Mi    | 24 Mi    |
| kube-system/network-problem-detector-pod-8svt5                              | 29 Mi    | 29 Mi   | 29 Mi    | 29 Mi    | 29 Mi    | 30 Mi    | 29 Mi    |
| kube-system/network-problem-detector-pod-c2wdm                              | 20 Mi    | 20 Mi   | 21 Mi    | 21 Mi    | 22 Mi    | 22 Mi    | 23 Mi    |
| kube-system/network-problem-detector-pod-c75rc                              | 24 Mi    | 25 Mi   | 25 Mi    | 25 Mi    | 23 Mi    | 24 Mi    | 24 Mi    |
| kube-system/network-problem-detector-pod-f4plc                              | 16 Mi    | 18 Mi   | 19 Mi    | 19 Mi    | 19 Mi    | 20 Mi    | 20 Mi    |
| kube-system/network-problem-detector-pod-hcxsv                              | 24 Mi    | 24 Mi   | 24 Mi    | 24 Mi    | 25 Mi    | 25 Mi    | 24 Mi    |
| kube-system/network-problem-detector-pod-hhbqr                              | 20 Mi    | 20 Mi   | 20 Mi    | 20 Mi    | 21 Mi    | 21 Mi    | 22 Mi    |
| kube-system/network-problem-detector-pod-p44l8                              | 33 Mi    | 33 Mi   | 33 Mi    | 34 Mi    | 34 Mi    | 34 Mi    | 34 Mi    |
| kube-system/node-exporter-cmgn6                                             | 15 Mi    | 15 Mi   | 15 Mi    | 15 Mi    | 15 Mi    | 15 Mi    | 15 Mi    |
| kube-system/node-exporter-hbqp6                                             | 23 Mi    | 23 Mi   | 23 Mi    | 23 Mi    | 23 Mi    | 23 Mi    | 23 Mi    |
| kube-system/node-exporter-lgpn7                                             | 14 Mi    | 14 Mi   | 13 Mi    | 14 Mi    | 14 Mi    | 14 Mi    | 14 Mi    |
| kube-system/node-exporter-lzz27                                             | 15 Mi    | 16 Mi   | 16 Mi    | 16 Mi    | 16 Mi    | 16 Mi    | 16 Mi    |
| kube-system/node-exporter-mmqph                                             | 20 Mi    | 21 Mi   | 21 Mi    | 21 Mi    | 21 Mi    | 21 Mi    | 21 Mi    |
| kube-system/node-exporter-qsdhl                                             | 10 Mi    | 10 Mi   | 10 Mi    | 10 Mi    | 10 Mi    | 11 Mi    | 10 Mi    |
| kube-system/node-exporter-sc76n                                             | 17 Mi    | 17 Mi   | 17 Mi    | 17 Mi    | 18 Mi    | 18 Mi    | 18 Mi    |
| kube-system/node-local-dns-cpu-worker-0-6khd2                               | 22 Mi    | 21 Mi   | 20 Mi    | 22 Mi    | 22 Mi    | 22 Mi    | 22 Mi    |
| kube-system/node-local-dns-cpu-worker-0-6sqvw                               | 22 Mi    | 22 Mi   | 24 Mi    | 24 Mi    | 25 Mi    | 25 Mi    | 25 Mi    |
| kube-system/node-local-dns-cpu-worker-0-8mblj                               | 23 Mi    | 23 Mi   | 23 Mi    | 23 Mi    | 19 Mi    | 21 Mi    | 22 Mi    |
| kube-system/node-local-dns-cpu-worker-0-klplf                               | 30 Mi    | 31 Mi   | 31 Mi    | 32 Mi    | 33 Mi    | 33 Mi    | 32 Mi    |
| kube-system/node-local-dns-cpu-worker-0-knndk                               | 26 Mi    | 26 Mi   | 25 Mi    | 26 Mi    | 24 Mi    | 25 Mi    | 24 Mi    |
| kube-system/node-local-dns-cpu-worker-0-nx6ht                               | 20 Mi    | 20 Mi   | 21 Mi    | 21 Mi    | 21 Mi    | 22 Mi    | 22 Mi    |
| kube-system/node-local-dns-cpu-worker-0-r2rzc                               | 17 Mi    | 18 Mi   | 18 Mi    | 20 Mi    | 20 Mi    | 20 Mi    | 21 Mi    |
| kube-system/node-problem-detector-74zcq                                     | 26 Mi    | 27 Mi   | 31 Mi    | 31 Mi    | 29 Mi    | 30 Mi    | 38 Mi    |
| kube-system/node-problem-detector-cdqdq                                     | 20 Mi    | 26 Mi   | 28 Mi    | 29 Mi    | 35 Mi    | 34 Mi    | 40 Mi    |
| kube-system/node-problem-detector-cgxjl                                     | 35 Mi    | 35 Mi   | 35 Mi    | 35 Mi    | 47 Mi    | 46 Mi    | 28 Mi    |
| kube-system/node-problem-detector-jk7dp                                     | 23 Mi    | 23 Mi   | 23 Mi    | 23 Mi    | 23 Mi    | 23 Mi    | 26 Mi    |
| kube-system/node-problem-detector-ls9d9                                     | 27 Mi    | 28 Mi   | 28 Mi    | 28 Mi    | 28 Mi    | 37 Mi    | 35 Mi    |
| kube-system/node-problem-detector-phn5t                                     | 42 Mi    | 42 Mi   | 42 Mi    | 43 Mi    | 43 Mi    | 53 Mi    | 43 Mi    |
| kube-system/node-problem-detector-w5v2s                                     | 20 Mi    | 21 Mi   | 21 Mi    | 21 Mi    | 21 Mi    | 28 Mi    | 34 Mi    |
| kube-system/vpn-shoot-0                                                     | 19 Mi    | 19 Mi   | 19 Mi    | 19 Mi    | 19 Mi    | 19 Mi    | 18 Mi    |
| kube-system/vpn-shoot-1                                                     | 13 Mi    | 13 Mi   | 13 Mi    | 13 Mi    | 13 Mi    | 13 Mi    | 13 Mi    |
| kwok-system/kwok-controller-d5dbcc9d9-2n4v5                                 | 23 Mi    | 26 Mi   | 34 Mi    | 38 Mi    | 43 Mi    | 57 Mi    | 86 Mi    |
| kyma-system/api-gateway-controller-manager-66d9d4d67d-f867w                 | 73 Mi    | 195 Mi  | 370 Mi   | 574 Mi   | 873 Mi   | 1258 Mi  | 2070 Mi  |
| kyma-system/application-connector-controller-manager-66f7775d77-wttlq       | 200 Mi   | 200 Mi  | 201 Mi   | 207 Mi   | 199 Mi   | 192 Mi   | 196 Mi   |
| kyma-system/btp-manager-controller-manager-7c6d47d7b6-dgbld                 | 31 Mi    | 31 Mi   | 32 Mi    | 31 Mi    | 30 Mi    | 33 Mi    | 33 Mi    |
| kyma-system/central-application-connectivity-validator-87c7b7947-m49hk      | 98 Mi    | 98 Mi   | 97 Mi    | 97 Mi    | 99 Mi    | 99 Mi    | 94 Mi    |
| kyma-system/central-application-connectivity-validator-87c7b7947-rrc7t      | 79 Mi    | 79 Mi   | 76 Mi    | 76 Mi    | 80 Mi    | 76 Mi    | 76 Mi    |
| kyma-system/central-application-gateway-7b695b9b57-jc94f                    | 80 Mi    | 80 Mi   | 76 Mi    | 76 Mi    | 75 Mi    | 74 Mi    | 77 Mi    |
| kyma-system/central-application-gateway-7b695b9b57-tws2l                    | 84 Mi    | 84 Mi   | 75 Mi    | 76 Mi    | 76 Mi    | 80 Mi    | 78 Mi    |
| kyma-system/compass-runtime-agent-567c6fbdb7-xjfv4                          | 88 Mi    | 89 Mi   | 90 Mi    | 91 Mi    | 93 Mi    | 95 Mi    | 92 Mi    |
| kyma-system/connectivity-proxy-0                                            | 207 Mi   | 207 Mi  | 203 Mi   | 203 Mi   | 208 Mi   | 208 Mi   | 206 Mi   |
| kyma-system/connectivity-proxy-operator-56c68f65c5-bpdj7                    | 392 Mi   | 392 Mi  | 389 Mi   | 394 Mi   | 394 Mi   | 392 Mi   | 391 Mi   |
| kyma-system/connectivity-proxy-region-configurations-controller-b58ff5wvgx2 | 233 Mi   | 233 Mi  | 226 Mi   | 227 Mi   | 232 Mi   | 227 Mi   | 230 Mi   |
| kyma-system/connectivity-proxy-restart-watcher-5656fc5865-gw495             | 361 Mi   | 361 Mi  | 355 Mi   | 356 Mi   | 356 Mi   | 361 Mi   | 358 Mi   |
| kyma-system/connectivity-proxy-sm-operator-59dff4d5df-2rn9b                 | 220 Mi   | 220 Mi  | 220 Mi   | 220 Mi   | 220 Mi   | 220 Mi   | 220 Mi   |
| kyma-system/eventing-nats-0                                                 | 43 Mi    | 44 Mi   | 44 Mi    | 45 Mi    | 45 Mi    | 45 Mi    | 45 Mi    |
| kyma-system/eventing-nats-1                                                 | 73 Mi    | 73 Mi   | 75 Mi    | 73 Mi    | 74 Mi    | 73 Mi    | 73 Mi    |
| kyma-system/eventing-nats-2                                                 | 53 Mi    | 53 Mi   | 53 Mi    | 52 Mi    | 53 Mi    | 54 Mi    | 54 Mi    |
| kyma-system/istio-controller-manager-6b7f78d558-fmlph                       | 63 Mi    | 184 Mi  | 387 Mi   | 633 Mi   | 666 Mi   | 1345 Mi  | 1345 Mi  |
| kyma-system/keda-admission-webhooks-855f88bf8-4c5kn                         | 12 Mi    | 12 Mi   | 12 Mi    | 12 Mi    | 12 Mi    | 12 Mi    | 12 Mi    |
| kyma-system/keda-manager-66d9959c47-mgjdf                                   | 68 Mi    | 68 Mi   | 69 Mi    | 69 Mi    | 68 Mi    | 68 Mi    | 68 Mi    |
| kyma-system/keda-operator-67b88dbf8b-pkgsg                                  | 50 Mi    | 50 Mi   | 50 Mi    | 50 Mi    | 50 Mi    | 50 Mi    | 51 Mi    |
| kyma-system/keda-operator-metrics-apiserver-6cbf7567fc-877pj                | 48 Mi    | 48 Mi   | 48 Mi    | 48 Mi    | 48 Mi    | 48 Mi    | 48 Mi    |
| kyma-system/kim-snatch-controller-manager-84cdb5df56-nwfp9                  | 21 Mi    | 21 Mi   | 21 Mi    | 21 Mi    | 19 Mi    | 20 Mi    | 21 Mi    |
| kyma-system/nats-manager-785744947f-w6rcb                                   | 91 Mi    | 91 Mi   | 85 Mi    | 85 Mi    | 85 Mi    | 88 Mi    | 86 Mi    |
| kyma-system/rma-kube-state-metrics-6879cd7599-q9254                         | 1827 Mi  | 1762 Mi | 1952 Mi  | 2826 Mi  | 3206 Mi  | 4146 Mi  | 4179 Mi  |
| kyma-system/rma-system-logs-agent-48hwt                                     | 28 Mi    | 29 Mi   | 31 Mi    | 30 Mi    | 30 Mi    | 30 Mi    | 30 Mi    |
| kyma-system/rma-system-logs-agent-4khk5                                     | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | 30 Mi    | 30 Mi    | 30 Mi    |
| kyma-system/rma-system-logs-agent-9pwkx                                     | 36 Mi    | 36 Mi   | 36 Mi    | 35 Mi    | 36 Mi    | 37 Mi    | 35 Mi    |
| kyma-system/rma-system-logs-agent-c8vmd                                     | 33 Mi    | 33 Mi   | 34 Mi    | 34 Mi    | 34 Mi    | 34 Mi    | 34 Mi    |
| kyma-system/rma-system-logs-agent-hxw25                                     | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | 34 Mi    | 34 Mi    | 35 Mi    |
| kyma-system/rma-system-logs-agent-vvnlm                                     | 35 Mi    | 36 Mi   | 36 Mi    | 36 Mi    | 35 Mi    | 37 Mi    | 37 Mi    |
| kyma-system/rma-system-logs-agent-xzgr7                                     | 37 Mi    | 37 Mi   | 37 Mi    | 39 Mi    | 37 Mi    | 38 Mi    | 39 Mi    |
| kyma-system/rma-system-logs-collector-565d78d588-vxmsv                      | 66 Mi    | 66 Mi   | 67 Mi    | 68 Mi    | 70 Mi    | 70 Mi    | 66 Mi    |
| kyma-system/rma-victoria-metrics-agent-77ff8d586c-h5zzp                     | 209 Mi   | 388 Mi  | 604 Mi   | 1029 Mi  | 1366 Mi  | 1832 Mi  | 1750 Mi  |
| kyma-system/sap-btp-operator-controller-manager-74bd4cbdc6-7mk86            | 42 Mi    | 42 Mi   | 42 Mi    | 42 Mi    | 42 Mi    | 42 Mi    | 42 Mi    |
| kyma-system/serverless-ctrl-mngr-7d4bf94d45-7gt2f                           | 21 Mi    | 23 Mi   | 23 Mi    | 23 Mi    | 23 Mi    | 23 Mi    | 23 Mi    |
| kyma-system/serverless-operator-b5c797f8c-v7jtd                             | 34 Mi    | 34 Mi   | 34 Mi    | 34 Mi    | 34 Mi    | 34 Mi    | 34 Mi    |
| kyma-system/skr-webhook-5444569bf6-bc4ph                                    | -        | -       | -        | -        | -        | -        | 7 Mi     |
| kyma-system/skr-webhook-5444569bf6-f27hp                                    | -        | -       | -        | 4 Mi     | 9 Mi     | 9 Mi     | 9 Mi     |
| kyma-system/skr-webhook-5444569bf6-fptvc                                    | -        | -       | -        | 4 Mi     | -        | -        | -        |
| kyma-system/skr-webhook-5444569bf6-k6b9q                                    | 9 Mi     | 11 Mi   | 11 Mi    | 12 Mi    | -        | -        | -        |
| kyma-system/skr-webhook-5444569bf6-lz8x4                                    | -        | -       | -        | -        | -        | -        | 9 Mi     |
| kyma-system/skr-webhook-5444569bf6-m7wcg                                    | -        | -       | -        | 6 Mi     | -        | -        | -        |
| kyma-system/telemetry-manager-674b746498-zzzt8                              | 46 Mi    | 45 Mi   | 46 Mi    | 46 Mi    | 43 Mi    | 45 Mi    | 45 Mi    |
| kyma-system/telemetry-metric-agent-4jlhk                                    | 245 Mi   | 245 Mi  | 264 Mi   | 277 Mi   | 282 Mi   | 314 Mi   | 326 Mi   |
| kyma-system/telemetry-metric-agent-d2rp9                                    | 111 Mi   | 122 Mi  | 128 Mi   | 163 Mi   | 178 Mi   | 211 Mi   | 344 Mi   |
| kyma-system/telemetry-metric-agent-dxrk5                                    | 249 Mi   | 256 Mi  | 261 Mi   | 281 Mi   | 284 Mi   | 313 Mi   | 328 Mi   |
| kyma-system/telemetry-metric-agent-nmzln                                    | 232 Mi   | 238 Mi  | 236 Mi   | 258 Mi   | 271 Mi   | 283 Mi   | 335 Mi   |
| kyma-system/telemetry-metric-agent-pf658                                    | 228 Mi   | 229 Mi  | 245 Mi   | 257 Mi   | 267 Mi   | 282 Mi   | 378 Mi   |
| kyma-system/telemetry-metric-agent-s8dvv                                    | 98 Mi    | 111 Mi  | 129 Mi   | 138 Mi   | 168 Mi   | 204 Mi   | 297 Mi   |
| kyma-system/telemetry-metric-agent-v5dk9                                    | 278 Mi   | 280 Mi  | 293 Mi   | 308 Mi   | 308 Mi   | 334 Mi   | 359 Mi   |
| kyma-system/telemetry-metric-gateway-f478fb4c5-t2h2g                        | 208 Mi   | 208 Mi  | 220 Mi   | 228 Mi   | 240 Mi   | 266 Mi   | 294 Mi   |
| kyma-system/telemetry-metric-gateway-f478fb4c5-t9sr9                        | 205 Mi   | 208 Mi  | 216 Mi   | 229 Mi   | 245 Mi   | 257 Mi   | 334 Mi   |
| kyma-system/telemetry-self-monitor-5994d4c965-jq8np                         | 66 Mi    | 66 Mi   | 64 Mi    | 64 Mi    | 87 Mi    | 87 Mi    | 66 Mi    |
| kyma-system/warden-admission-767f55476-bzbfg                                | 29 Mi    | 29 Mi   | 27 Mi    | 31 Mi    | 30 Mi    | 29 Mi    | 32 Mi    |
| kyma-system/warden-operator-69bf7ff597-rv5bv                                | 43 Mi    | 157 Mi  | 314 Mi   | 574 Mi   | 685 Mi   | 1015 Mi  | 1015 Mi  |
| movies-rest/kwiatekus-movies-rest-58f4876589-wngbj                          | 244 Mi   | 244 Mi  | 240 Mi   | 240 Mi   | 242 Mi   | 242 Mi   | 238 Mi   |
| nats-jetstream-demo/jetstream-consumer-54bb874699-285b2                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2mwpp                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2nnpk                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2qbmd                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2sl5f                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2xskb                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-45hdq                     | -        | 8 Mi    | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-4lplf                     | -        | 6 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-5qrjp                     | -        | -       | -        | -        | 6 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-5zwpv                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-675nn                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-6b2rc                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-6brnf                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-6k4bn                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-72z2z                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-78xwt                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-7fxdf                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-7sdg8                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-7vjd8                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-82f2n                     | -        | -       | -        | 7 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-89v5j                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-8tvr4                     | -        | -       | 8 Mi     | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-8wtbl                     | -        | -       | -        | -        | -        | -        | 7 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-95ttw                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-97lv2                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9g2sj                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9kndk                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9rjsk                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9rp8n                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9z6p2                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-b6mwz                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-b776q                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-bgdjq                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-bv2wx                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-c4zwv                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-c6b92                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-cjp9l                     | -        | -       | -        | -        | -        | 7 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-ckqn4                     | -        | -       | -        | -        | -        | 7 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-ctrkv                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-d5ldn                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-d8h2g                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-ddk5g                     | -        | -       | -        | -        | 7 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-dh5cc                     | -        | 6 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-f26kj                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-fpnw6                     | -        | -       | -        | -        | -        | -        | 7 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-fpvxc                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-fwndb                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-g7sst                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-g9qkx                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-gb4s4                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-gcvtl                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-gpd8k                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-gpt9n                     | 8 Mi     | -       | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-grtfc                     | -        | -       | -        | -        | -        | 7 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-gssnn                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-hnzf8                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-hrc8x                     | -        | -       | -        | -        | -        | 6 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-htlc9                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-j2mc4                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-jlpnr                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-jvjww                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-khcq9                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-kst6d                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-l8457                     | 8 Mi     | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-lxbns                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-m7x4g                     | -        | -       | -        | -        | -        | -        | 7 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-mql2s                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-n4z4t                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-n5kpk                     | -        | -       | -        | -        | -        | 7 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-n69kr                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-nmm49                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-nvxj2                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-pfb2f                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-q4sph                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-qf29f                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-qj7ph                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-qspnh                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-qxq4t                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-rhdmk                     | -        | -       | -        | -        | -        | 7 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-rwsts                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-s9xdd                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-sb2k5                     | -        | -       | -        | -        | -        | -        | 6 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-skxjh                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-slqcm                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-sv89g                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-svtmq                     | -        | -       | -        | -        | -        | -        | 6 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-t25k2                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-t5c9m                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-t5tj2                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-tcv2t                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-tdx2b                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-tf4fd                     | -        | -       | -        | -        | -        | -        | 7 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-tkhpf                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-tl288                     | -        | -       | -        | 7 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-v9s6m                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-vnc4k                     | -        | -       | -        | -        | -        | -        | 6 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-vwdn5                     | -        | -       | -        | -        | -        | -        | 6 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-w59jj                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-w98f7                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-wr6jl                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-wsvt6                     | -        | -       | -        | -        | -        | -        | 8 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-wvrfw                     | 8 Mi     | -       | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-x4q6j                     | -        | -       | -        | -        | -        | -        | 6 Mi     |
| nats-jetstream-demo/jetstream-consumer-54bb874699-xwzgr                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-xxw8l                     | -        | -       | -        | -        | 7 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-xznjr                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-z49gd                     | -        | -       | -        | -        | -        | 8 Mi     | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-zljdw                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-producer-86547dcdfd-bfcdk                     | 2 Mi     | 2 Mi    | 2 Mi     | 2 Mi     | 2 Mi     | 7 Mi     | 2 Mi     |
| registry-proxy/registry-proxy-controller-5955cc9f9f-g82cd                   | 86 Mi    | 127 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| registry-proxy/registry-proxy-operator-5f8dc4d9b7-z8mqb                     | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | 33 Mi    | 33 Mi    | 33 Mi    |
| sap-transp-proxy-system/sap-transp-proxy-healthcheck-58966dd54b-znt8j       | 70 Mi    | 70 Mi   | 70 Mi    | 71 Mi    | 75 Mi    | 75 Mi    | 79 Mi    |
| sap-transp-proxy-system/sap-transp-proxy-manager-64586c989b-nffp5           | 73 Mi    | 73 Mi   | 69 Mi    | 72 Mi    | 75 Mi    | 75 Mi    | 78 Mi    |
| sap-transp-proxy-system/sap-transp-proxy-operator-76d5c8fd84-9zmcd          | 98 Mi    | 95 Mi   | 95 Mi    | 121 Mi   | 121 Mi   | 123 Mi   | 131 Mi   |
| ztis-agent-system/ztis-operator-controller-manager-c89c75d56-dvldl          | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | 31 Mi    | 31 Mi    | 31 Mi    |