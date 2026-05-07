# Secrets

## Summary

Pods with memory growth exceeding 200 Mi from baseline to 3000 resources (last measured column; pods that failed before 3000 res show delta at last valid measurement):

| pod (namespace/name)                                                        | baseline | 500 res | 1000 res | 2000 res | 3000 res | 4000 res | 5000 res | delta    |
| --------------------------------------------------------------------------- | -------- | ------- | -------- | -------- | -------- | -------- | -------- | -------- |
| istio-system/istiod-66b4585c8b-phjqw                                        | 151 Mi   | 457 Mi  | 883 Mi   | 1087 Mi  | 1859 Mi  | -        | FAILING  | +1708 Mi |
| istio-system/istiod-66b4585c8b-slnvr                                        | 102 Mi   | 398 Mi  | 945 Mi   | 1322 Mi  | 1664 Mi  | -        | FAILING  | +1562 Mi |
| korifi/korifi-controllers-controller-manager-5ddd86668-85tk6                | 965 Mi   | 991 Mi  | 1389 Mi  | 1560 Mi  | 2281 Mi  | -        | -        | +1316 Mi |
| kyma-system/keda-operator-67b88dbf8b-pkgsg                                  | 50 Mi    | 344 Mi  | 690 Mi   | 768 Mi   | FAILING  | FAILING  | FAILING  | +718 Mi  |
| docker-registry/dockerregistry-operator-587f69dbc4-rtz9z                    | 77 Mi    | 379 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  | +302 Mi  |
| cap-operator-system/cap-operator-controller-9c9bc97bb-5kpsc                 | 27 Mi    | 298 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  | +271 Mi  |
| kyma-system/application-connector-controller-manager-66f7775d77-wttlq       | 200 Mi   | 451 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  | +251 Mi  |
| cfapi-system/cfapi-operator-989bb945-4mn9s                                  | 202 Mi   | 404 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | -        | +202 Mi  |

Pods that have a valid baseline but start failing in further steps:

| pod (namespace/name)                                                        | baseline | 500 res | 1000 res | 2000 res | 3000 res | 4000 res | 5000 res |
| --------------------------------------------------------------------------- | -------- | ------- | -------- | -------- | -------- | -------- | -------- |
| cfapi-system/contour-contour-57fd8d96b8-mn452                               | 146 Mi   | FAILING | FAILING  | FAILING  | FAILING  | -        | FAILING  |
| registry-proxy/registry-proxy-operator-5f8dc4d9b7-z8mqb                     | 46 Mi    | FAILING | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| cap-operator-system/cap-operator-controller-9c9bc97bb-5kpsc                 | 27 Mi    | 298 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| docker-registry/dockerregistry-operator-587f69dbc4-rtz9z                    | 77 Mi    | 379 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| kyma-system/application-connector-controller-manager-66f7775d77-wttlq       | 200 Mi   | 451 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| cfapi-system/cfapi-operator-989bb945-4mn9s                                  | 202 Mi   | 404 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | -        |
| kyma-system/keda-operator-67b88dbf8b-pkgsg                                  | 50 Mi    | 344 Mi  | 690 Mi   | 768 Mi   | FAILING  | FAILING  | FAILING  |
| istio-system/istiod-66b4585c8b-phjqw                                        | 151 Mi   | 457 Mi  | 883 Mi   | 1087 Mi  | 1859 Mi  | -        | FAILING  |
| istio-system/istiod-66b4585c8b-slnvr                                        | 102 Mi   | 398 Mi  | 945 Mi   | 1322 Mi  | 1664 Mi  | -        | FAILING  |

## Full report

| pod (namespace/name)                                                        | baseline | 500 res | 1000 res | 2000 res | 3000 res | 4000 res | 5000 res |
| --------------------------------------------------------------------------- | -------- | ------- | -------- | -------- | -------- | -------- | -------- |
| antek-system/fake-metrics-746484997-jncvm                                   | 10 Mi    | 10 Mi   | 10 Mi    | 10 Mi    | 10 Mi    | -        | -        |
| antek-system/opensearch-external-scaler-fdbb7ffc-97zll                      | 8 Mi     | 8 Mi    | 8 Mi     | 8 Mi     | 8 Mi     | -        | -        |
| cap-operator-system/cap-operator-controller-9c9bc97bb-5kpsc                 | 27 Mi    | 298 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| cap-operator-system/cap-operator-controller-manager-57cdf94999-dspsd        | 46 Mi    | 52 Mi   | 54 Mi    | 64 Mi    | 65 Mi    | -        | -        |
| cap-operator-system/cap-operator-subscription-server-d997f7d8f-frq6k        | 9 Mi     | 9 Mi    | 9 Mi     | 9 Mi     | 9 Mi     | -        | -        |
| cap-operator-system/cap-operator-webhook-786459c7ff-kl8pm                   | 11 Mi    | 11 Mi   | 11 Mi    | 11 Mi    | 11 Mi    | -        | -        |
| cfapi-system/btp-service-broker-cc965b778-qjjpz                             | 15 Mi    | 15 Mi   | 15 Mi    | 15 Mi    | 15 Mi    | -        | -        |
| cfapi-system/cfapi-operator-989bb945-4mn9s                                  | 202 Mi   | 404 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | -        |
| cfapi-system/contour-contour-57fd8d96b8-mn452                               | 146 Mi   | FAILING | FAILING  | FAILING  | FAILING  | -        | FAILING  |
| cfapi-system/contour-envoy-2v5vq                                            | 47 Mi    | 48 Mi   | 48 Mi    | 50 Mi    | 50 Mi    | -        | -        |
| cfapi-system/contour-envoy-7x65v                                            | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | 33 Mi    | -        | -        |
| cfapi-system/contour-envoy-dw2rn                                            | 35 Mi    | 35 Mi   | 35 Mi    | 35 Mi    | 36 Mi    | -        | -        |
| cfapi-system/contour-envoy-jmtx6                                            | 31 Mi    | 31 Mi   | 31 Mi    | 31 Mi    | -        | -        | -        |
| cfapi-system/contour-envoy-vpbpx                                            | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | 33 Mi    | -        | -        |
| cfapi-system/contour-envoy-xbjgh                                            | 34 Mi    | 34 Mi   | 34 Mi    | 34 Mi    | 34 Mi    | -        | -        |
| cfapi-system/contour-envoy-zdqns                                            | 35 Mi    | 35 Mi   | 35 Mi    | 35 Mi    | 35 Mi    | -        | -        |
| demo-app/http-echo-677d479d69-s24pf                                         | 73 Mi    | 77 Mi   | 77 Mi    | 76 Mi    | 75 Mi    | -        | -        |
| dev/api-postgresql-go-67f4797b5c-wppkk                                      | 72 Mi    | 72 Mi   | 72 Mi    | 76 Mi    | 73 Mi    | -        | -        |
| dev/commerce-mock-75648c645b-lh26b                                          | 181 Mi   | 181 Mi  | 181 Mi   | 183 Mi   | 182 Mi   | -        | -        |
| dev/fe-ui5-postgresql-55ff69fbd7-dhqgm                                      | 68 Mi    | 70 Mi   | 67 Mi    | 71 Mi    | 70 Mi    | -        | -        |
| docker-registry/dockerregistry-fd5f6685c-tfxws                              | 86 Mi    | 86 Mi   | 87 Mi    | 87 Mi    | 87 Mi    | -        | -        |
| docker-registry/dockerregistry-operator-587f69dbc4-rtz9z                    | 77 Mi    | 379 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| istio-system/istio-cni-node-2fll2                                           | 52 Mi    | 52 Mi   | 52 Mi    | 52 Mi    | 52 Mi    | -        | -        |
| istio-system/istio-cni-node-gsf2b                                           | 36 Mi    | 36 Mi   | 36 Mi    | 37 Mi    | 36 Mi    | -        | -        |
| istio-system/istio-cni-node-h26x8                                           | 29 Mi    | 29 Mi   | 29 Mi    | 29 Mi    | 30 Mi    | -        | -        |
| istio-system/istio-cni-node-httjn                                           | 30 Mi    | 30 Mi   | 30 Mi    | 33 Mi    | 36 Mi    | -        | -        |
| istio-system/istio-cni-node-lkpqw                                           | 21 Mi    | 21 Mi   | 21 Mi    | 20 Mi    | -        | -        | -        |
| istio-system/istio-cni-node-p6qxd                                           | 32 Mi    | 32 Mi   | 29 Mi    | 30 Mi    | 31 Mi    | -        | -        |
| istio-system/istio-cni-node-vzr4g                                           | 22 Mi    | 23 Mi   | 23 Mi    | 23 Mi    | 24 Mi    | -        | -        |
| istio-system/istio-cni-node-wsspg                                           | 48 Mi    | 48 Mi   | 48 Mi    | 49 Mi    | 48 Mi    | -        | -        |
| istio-system/istio-ingressgateway-76974d6f7f-65kbd                          | 93 Mi    | 93 Mi   | 98 Mi    | 99 Mi    | 101 Mi   | -        | -        |
| istio-system/istio-ingressgateway-76974d6f7f-c5h5v                          | 74 Mi    | 74 Mi   | 77 Mi    | 79 Mi    | 77 Mi    | -        | -        |
| istio-system/istio-ingressgateway-76974d6f7f-thkxr                          | 154 Mi   | 154 Mi  | 154 Mi   | 152 Mi   | 155 Mi   | -        | -        |
| istio-system/istiod-66b4585c8b-phjqw                                        | 151 Mi   | 457 Mi  | 883 Mi   | 1087 Mi  | 1859 Mi  | -        | FAILING  |
| istio-system/istiod-66b4585c8b-slnvr                                        | 102 Mi   | 398 Mi  | 945 Mi   | 1322 Mi  | 1664 Mi  | -        | FAILING  |
| korifi/korifi-api-deployment-6d8c9f784d-lr6kh                               | 17 Mi    | 17 Mi   | 17 Mi    | 17 Mi    | 17 Mi    | -        | -        |
| korifi/korifi-controllers-controller-manager-5ddd86668-85tk6                | 965 Mi   | 991 Mi  | 1389 Mi  | 1560 Mi  | 2281 Mi  | -        | -        |
| kpack/kpack-controller-5c6866cf46-twcx6                                     | 53 Mi    | 53 Mi   | 53 Mi    | 53 Mi    | 53 Mi    | -        | -        |
| kpack/kpack-webhook-54cf8f75bf-pk754                                        | 80 Mi    | 80 Mi   | 89 Mi    | 103 Mi   | 100 Mi   | -        | -        |
| kube-system/apiserver-proxy-479wm                                           | 31 Mi    | 31 Mi   | 31 Mi    | 31 Mi    | 31 Mi    | -        | -        |
| kube-system/apiserver-proxy-5lr8p                                           | 41 Mi    | 41 Mi   | 41 Mi    | 41 Mi    | 42 Mi    | -        | -        |
| kube-system/apiserver-proxy-7skgj                                           | 25 Mi    | 25 Mi   | 25 Mi    | 25 Mi    | -        | -        | -        |
| kube-system/apiserver-proxy-8kw9m                                           | 32 Mi    | 32 Mi   | 32 Mi    | 32 Mi    | 32 Mi    | -        | -        |
| kube-system/apiserver-proxy-997dr                                           | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | 30 Mi    | -        | -        |
| kube-system/apiserver-proxy-hf6qf                                           | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | 30 Mi    | -        | -        |
| kube-system/apiserver-proxy-jpxjc                                           | 52 Mi    | 52 Mi   | 52 Mi    | 52 Mi    | 52 Mi    | -        | -        |
| kube-system/apiserver-proxy-q688t                                           | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | 33 Mi    | -        | -        |
| kube-system/blackbox-exporter-59694df6-tvrdm                                | 17 Mi    | 17 Mi   | 17 Mi    | 17 Mi    | 17 Mi    | -        | -        |
| kube-system/blackbox-exporter-59694df6-vwn7t                                | 13 Mi    | 13 Mi   | 15 Mi    | 15 Mi    | 16 Mi    | -        | -        |
| kube-system/calico-node-5dd2h                                               | 118 Mi   | 119 Mi  | 119 Mi   | 122 Mi   | -        | -        | -        |
| kube-system/calico-node-gw2x6                                               | 214 Mi   | 212 Mi  | 215 Mi   | 215 Mi   | 216 Mi   | -        | -        |
| kube-system/calico-node-jg47c                                               | 204 Mi   | 204 Mi  | 205 Mi   | 224 Mi   | 205 Mi   | -        | -        |
| kube-system/calico-node-kzvcd                                               | 208 Mi   | 207 Mi  | 207 Mi   | 207 Mi   | 229 Mi   | -        | -        |
| kube-system/calico-node-lmvlz                                               | 215 Mi   | 216 Mi  | 215 Mi   | 218 Mi   | 228 Mi   | -        | -        |
| kube-system/calico-node-ntsb5                                               | 189 Mi   | 190 Mi  | 190 Mi   | 216 Mi   | 216 Mi   | -        | -        |
| kube-system/calico-node-vertical-autoscaler-57d6f4f8f-t6p6x                 | 12 Mi    | 12 Mi   | 13 Mi    | 13 Mi    | 13 Mi    | -        | -        |
| kube-system/calico-node-vjp5x                                               | 230 Mi   | 232 Mi  | 232 Mi   | 255 Mi   | 251 Mi   | -        | -        |
| kube-system/calico-node-xfhnq                                               | 214 Mi   | 199 Mi  | 199 Mi   | 200 Mi   | 221 Mi   | -        | -        |
| kube-system/calico-typha-deploy-bd76c7b67-496gm                             | 102 Mi   | 102 Mi  | 101 Mi   | 102 Mi   | 104 Mi   | -        | -        |
| kube-system/calico-typha-deploy-bd76c7b67-5k8hj                             | 97 Mi    | 97 Mi   | 97 Mi    | 98 Mi    | 96 Mi    | -        | -        |
| kube-system/calico-typha-horizontal-autoscaler-66b75b5d95-rzspb             | 13 Mi    | 13 Mi   | 13 Mi    | 13 Mi    | 13 Mi    | -        | -        |
| kube-system/calico-typha-vertical-autoscaler-54b5769b9-8fjp2                | 12 Mi    | 13 Mi   | 13 Mi    | 13 Mi    | 13 Mi    | -        | -        |
| kube-system/cloud-node-manager-bw6nb                                        | 44 Mi    | 44 Mi   | 44 Mi    | 44 Mi    | 44 Mi    | -        | -        |
| kube-system/cloud-node-manager-d4qkw                                        | 17 Mi    | 17 Mi   | 17 Mi    | 17 Mi    | 17 Mi    | -        | -        |
| kube-system/cloud-node-manager-gmgxt                                        | 40 Mi    | 40 Mi   | 40 Mi    | 40 Mi    | 40 Mi    | -        | -        |
| kube-system/cloud-node-manager-k25vj                                        | 22 Mi    | 22 Mi   | 22 Mi    | 22 Mi    | 22 Mi    | -        | -        |
| kube-system/cloud-node-manager-krnf2                                        | 15 Mi    | 15 Mi   | 15 Mi    | 15 Mi    | 15 Mi    | -        | -        |
| kube-system/cloud-node-manager-lrzq5                                        | 28 Mi    | 28 Mi   | 28 Mi    | 28 Mi    | 28 Mi    | -        | -        |
| kube-system/cloud-node-manager-nzjvr                                        | 16 Mi    | 16 Mi   | 16 Mi    | 16 Mi    | 16 Mi    | -        | -        |
| kube-system/cloud-node-manager-r9xkm                                        | 14 Mi    | 14 Mi   | 14 Mi    | 15 Mi    | -        | -        | -        |
| kube-system/coredns-57dd88d9c5-4mr4k                                        | 32 Mi    | 32 Mi   | 32 Mi    | 31 Mi    | 31 Mi    | -        | -        |
| kube-system/coredns-57dd88d9c5-986ls                                        | 23 Mi    | 23 Mi   | 24 Mi    | 24 Mi    | 23 Mi    | -        | -        |
| kube-system/csi-driver-node-disk-62qkz                                      | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | 30 Mi    | -        | -        |
| kube-system/csi-driver-node-disk-7vzk5                                      | 37 Mi    | 37 Mi   | 37 Mi    | 37 Mi    | 38 Mi    | -        | -        |
| kube-system/csi-driver-node-disk-8vsm8                                      | 29 Mi    | 29 Mi   | 29 Mi    | 29 Mi    | 29 Mi    | -        | -        |
| kube-system/csi-driver-node-disk-99dwq                                      | 28 Mi    | 28 Mi   | 28 Mi    | 28 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-disk-gxksp                                      | 37 Mi    | 37 Mi   | 37 Mi    | 37 Mi    | 37 Mi    | -        | -        |
| kube-system/csi-driver-node-disk-hqrj7                                      | 32 Mi    | 32 Mi   | 32 Mi    | 33 Mi    | 32 Mi    | -        | -        |
| kube-system/csi-driver-node-disk-q7qns                                      | 44 Mi    | 44 Mi   | 44 Mi    | 44 Mi    | 44 Mi    | -        | -        |
| kube-system/csi-driver-node-disk-rpkkb                                      | 42 Mi    | 42 Mi   | 42 Mi    | 45 Mi    | 45 Mi    | -        | -        |
| kube-system/csi-driver-node-file-2zh57                                      | 44 Mi    | 44 Mi   | 44 Mi    | 44 Mi    | 44 Mi    | -        | -        |
| kube-system/csi-driver-node-file-49s9r                                      | 27 Mi    | 27 Mi   | 27 Mi    | 27 Mi    | 27 Mi    | -        | -        |
| kube-system/csi-driver-node-file-4cqsk                                      | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | 30 Mi    | -        | -        |
| kube-system/csi-driver-node-file-7tk8r                                      | 42 Mi    | 42 Mi   | 42 Mi    | 45 Mi    | 45 Mi    | -        | -        |
| kube-system/csi-driver-node-file-gfbhk                                      | 37 Mi    | 37 Mi   | 37 Mi    | 38 Mi    | 38 Mi    | -        | -        |
| kube-system/csi-driver-node-file-gx842                                      | 28 Mi    | 28 Mi   | 28 Mi    | 28 Mi    | 28 Mi    | -        | -        |
| kube-system/csi-driver-node-file-j5nqx                                      | 40 Mi    | 40 Mi   | 40 Mi    | 40 Mi    | 40 Mi    | -        | -        |
| kube-system/csi-driver-node-file-jt68v                                      | 28 Mi    | 28 Mi   | 28 Mi    | 28 Mi    | -        | -        | -        |
| kube-system/egress-filter-applier-bcc7b                                     | 16 Mi    | 16 Mi   | 16 Mi    | 16 Mi    | 18 Mi    | -        | -        |
| kube-system/egress-filter-applier-ddz6s                                     | 12 Mi    | 12 Mi   | 12 Mi    | 12 Mi    | 14 Mi    | -        | -        |
| kube-system/egress-filter-applier-fcwfv                                     | 12 Mi    | 12 Mi   | 12 Mi    | 11 Mi    | 13 Mi    | -        | -        |
| kube-system/egress-filter-applier-k8bxb                                     | 12 Mi    | 12 Mi   | 12 Mi    | 12 Mi    | 12 Mi    | -        | -        |
| kube-system/egress-filter-applier-mpcwc                                     | 17 Mi    | 17 Mi   | 17 Mi    | 16 Mi    | 19 Mi    | -        | -        |
| kube-system/egress-filter-applier-q5zr5                                     | 16 Mi    | 16 Mi   | 16 Mi    | 16 Mi    | 16 Mi    | -        | -        |
| kube-system/egress-filter-applier-t4g7n                                     | 12 Mi    | 12 Mi   | 12 Mi    | 12 Mi    | 13 Mi    | -        | -        |
| kube-system/egress-filter-applier-t5z9f                                     | 9 Mi     | 9 Mi    | 9 Mi     | 9 Mi     | -        | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-2wczl                           | 22 Mi    | 24 Mi   | 25 Mi    | 25 Mi    | -        | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-69l5f                           | 29 Mi    | 29 Mi   | 29 Mi    | 29 Mi    | 32 Mi    | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-cbd4s                           | 39 Mi    | 40 Mi   | 39 Mi    | 39 Mi    | 40 Mi    | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-hrsbg                           | 47 Mi    | 48 Mi   | 49 Mi    | 51 Mi    | 52 Mi    | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-lx7tt                           | 35 Mi    | 35 Mi   | 34 Mi    | 34 Mi    | 35 Mi    | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-n25mv                           | 54 Mi    | 54 Mi   | 53 Mi    | 54 Mi    | 54 Mi    | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-nkmzf                           | 28 Mi    | 29 Mi   | 29 Mi    | 29 Mi    | 30 Mi    | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-qlhfv                           | 34 Mi    | 34 Mi   | 35 Mi    | 36 Mi    | 34 Mi    | -        | -        |
| kube-system/metrics-server-bd6684798-d57vf                                  | 82 Mi    | 83 Mi   | 83 Mi    | 84 Mi    | 84 Mi    | -        | -        |
| kube-system/metrics-server-bd6684798-psls2                                  | 335 Mi   | 337 Mi  | 337 Mi   | 336 Mi   | 337 Mi   | -        | -        |
| kube-system/network-problem-detector-host-52ntc                             | 19 Mi    | 19 Mi   | 20 Mi    | 20 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-host-58tnx                             | 22 Mi    | 22 Mi   | 23 Mi    | 24 Mi    | 24 Mi    | -        | -        |
| kube-system/network-problem-detector-host-bz8jp                             | 26 Mi    | 27 Mi   | 27 Mi    | 29 Mi    | 28 Mi    | -        | -        |
| kube-system/network-problem-detector-host-gbwfs                             | 28 Mi    | 28 Mi   | 28 Mi    | 27 Mi    | 28 Mi    | -        | -        |
| kube-system/network-problem-detector-host-l6t7s                             | 23 Mi    | 23 Mi   | 24 Mi    | 24 Mi    | 24 Mi    | -        | -        |
| kube-system/network-problem-detector-host-p8728                             | 23 Mi    | 24 Mi   | 24 Mi    | 25 Mi    | 25 Mi    | -        | -        |
| kube-system/network-problem-detector-host-ts5jd                             | 19 Mi    | 20 Mi   | 20 Mi    | 21 Mi    | 21 Mi    | -        | -        |
| kube-system/network-problem-detector-host-xx48m                             | 28 Mi    | 29 Mi   | 28 Mi    | 28 Mi    | 28 Mi    | -        | -        |
| kube-system/network-problem-detector-pod-8svt5                              | 27 Mi    | 28 Mi   | 28 Mi    | 27 Mi    | 27 Mi    | -        | -        |
| kube-system/network-problem-detector-pod-c2wdm                              | 21 Mi    | 21 Mi   | 22 Mi    | 22 Mi    | 23 Mi    | -        | -        |
| kube-system/network-problem-detector-pod-c75rc                              | 23 Mi    | 24 Mi   | 24 Mi    | 25 Mi    | 24 Mi    | -        | -        |
| kube-system/network-problem-detector-pod-f4plc                              | 19 Mi    | 20 Mi   | 20 Mi    | 22 Mi    | 22 Mi    | -        | -        |
| kube-system/network-problem-detector-pod-hcxsv                              | 40 Mi    | 41 Mi   | 42 Mi    | 44 Mi    | 43 Mi    | -        | -        |
| kube-system/network-problem-detector-pod-hhbqr                              | 27 Mi    | 26 Mi   | 26 Mi    | 27 Mi    | 27 Mi    | -        | -        |
| kube-system/network-problem-detector-pod-ncjqb                              | 20 Mi    | 20 Mi   | 21 Mi    | 22 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-pod-p44l8                              | 29 Mi    | 30 Mi   | 31 Mi    | 29 Mi    | 31 Mi    | -        | -        |
| kube-system/node-exporter-cmgn6                                             | 16 Mi    | 16 Mi   | 16 Mi    | 16 Mi    | 16 Mi    | -        | -        |
| kube-system/node-exporter-d5ksq                                             | 10 Mi    | 10 Mi   | 10 Mi    | 10 Mi    | -        | -        | -        |
| kube-system/node-exporter-hbqp6                                             | 22 Mi    | 23 Mi   | 23 Mi    | 23 Mi    | 24 Mi    | -        | -        |
| kube-system/node-exporter-lgpn7                                             | 25 Mi    | 25 Mi   | 25 Mi    | 25 Mi    | 25 Mi    | -        | -        |
| kube-system/node-exporter-lzz27                                             | 26 Mi    | 27 Mi   | 28 Mi    | 28 Mi    | 28 Mi    | -        | -        |
| kube-system/node-exporter-mmqph                                             | 20 Mi    | 20 Mi   | 20 Mi    | 20 Mi    | 21 Mi    | -        | -        |
| kube-system/node-exporter-qsdhl                                             | 10 Mi    | 10 Mi   | 11 Mi    | 11 Mi    | 11 Mi    | -        | -        |
| kube-system/node-exporter-sc76n                                             | 17 Mi    | 18 Mi   | 18 Mi    | 19 Mi    | 18 Mi    | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-6khd2                               | 34 Mi    | 35 Mi   | 36 Mi    | 38 Mi    | 38 Mi    | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-6sqvw                               | 25 Mi    | 24 Mi   | 24 Mi    | 26 Mi    | 26 Mi    | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-8mblj                               | 21 Mi    | 21 Mi   | 23 Mi    | 24 Mi    | 22 Mi    | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-klplf                               | 31 Mi    | 31 Mi   | 31 Mi    | 31 Mi    | 31 Mi    | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-knndk                               | 28 Mi    | 29 Mi   | 29 Mi    | 28 Mi    | 29 Mi    | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-nx6ht                               | 41 Mi    | 41 Mi   | 43 Mi    | 44 Mi    | 45 Mi    | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-r2rzc                               | 20 Mi    | 22 Mi   | 20 Mi    | 22 Mi    | 22 Mi    | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-v79rl                               | 18 Mi    | 18 Mi   | 18 Mi    | 20 Mi    | -        | -        | -        |
| kube-system/node-problem-detector-74zcq                                     | 27 Mi    | 39 Mi   | 32 Mi    | 45 Mi    | 34 Mi    | -        | -        |
| kube-system/node-problem-detector-cdqdq                                     | 27 Mi    | 35 Mi   | 35 Mi    | 31 Mi    | 39 Mi    | -        | -        |
| kube-system/node-problem-detector-cgxjl                                     | 32 Mi    | 32 Mi   | 33 Mi    | 41 Mi    | 40 Mi    | -        | -        |
| kube-system/node-problem-detector-jk7dp                                     | 70 Mi    | 78 Mi   | 75 Mi    | 79 Mi    | 94 Mi    | -        | -        |
| kube-system/node-problem-detector-ls9d9                                     | 60 Mi    | 60 Mi   | 65 Mi    | 71 Mi    | 66 Mi    | -        | -        |
| kube-system/node-problem-detector-phn5t                                     | 31 Mi    | 32 Mi   | 32 Mi    | 31 Mi    | 32 Mi    | -        | -        |
| kube-system/node-problem-detector-v98qz                                     | 20 Mi    | 21 Mi   | 21 Mi    | 21 Mi    | -        | -        | -        |
| kube-system/node-problem-detector-w5v2s                                     | 22 Mi    | 22 Mi   | 22 Mi    | 23 Mi    | 23 Mi    | -        | -        |
| kube-system/vpn-shoot-0                                                     | 17 Mi    | 17 Mi   | 17 Mi    | 17 Mi    | 17 Mi    | -        | -        |
| kube-system/vpn-shoot-1                                                     | 20 Mi    | 20 Mi   | 20 Mi    | 23 Mi    | 23 Mi    | -        | -        |
| kwok-system/kwok-controller-d5dbcc9d9-qvsfk                                 | 74 Mi    | 77 Mi   | 77 Mi    | 81 Mi    | 83 Mi    | -        | -        |
| kyma-system/api-gateway-controller-manager-66d9d4d67d-gw8lk                 | 69 Mi    | 69 Mi   | 69 Mi    | 70 Mi    | 75 Mi    | -        | -        |
| kyma-system/application-connector-controller-manager-66f7775d77-wttlq       | 200 Mi   | 451 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| kyma-system/btp-manager-controller-manager-7c6d47d7b6-dgbld                 | 35 Mi    | 37 Mi   | 39 Mi    | 39 Mi    | 38 Mi    | -        | -        |
| kyma-system/central-application-connectivity-validator-87c7b7947-m49hk      | 83 Mi    | 84 Mi   | 84 Mi    | 83 Mi    | 85 Mi    | -        | -        |
| kyma-system/central-application-connectivity-validator-87c7b7947-rrc7t      | 79 Mi    | 80 Mi   | 76 Mi    | 80 Mi    | 78 Mi    | -        | -        |
| kyma-system/central-application-gateway-7b695b9b57-jc94f                    | 110 Mi   | 110 Mi  | 110 Mi   | 110 Mi   | 110 Mi   | -        | -        |
| kyma-system/central-application-gateway-7b695b9b57-tws2l                    | 82 Mi    | 82 Mi   | 79 Mi    | 84 Mi    | 84 Mi    | -        | -        |
| kyma-system/compass-runtime-agent-567c6fbdb7-xjfv4                          | 107 Mi   | 108 Mi  | 109 Mi   | 109 Mi   | 110 Mi   | -        | -        |
| kyma-system/connectivity-proxy-0                                            | 185 Mi   | 186 Mi  | 195 Mi   | 195 Mi   | 192 Mi   | -        | -        |
| kyma-system/connectivity-proxy-operator-56c68f65c5-bpdj7                    | 395 Mi   | 397 Mi  | 397 Mi   | 400 Mi   | 400 Mi   | -        | -        |
| kyma-system/connectivity-proxy-region-configurations-controller-b58ff5wvgx2 | 197 Mi   | 197 Mi  | 198 Mi   | 200 Mi   | 205 Mi   | -        | -        |
| kyma-system/connectivity-proxy-restart-watcher-5656fc5865-gw495             | 294 Mi   | 294 Mi  | 299 Mi   | 300 Mi   | 303 Mi   | -        | -        |
| kyma-system/connectivity-proxy-sm-operator-59dff4d5df-2rn9b                 | 228 Mi   | 229 Mi  | 233 Mi   | 235 Mi   | 235 Mi   | -        | -        |
| kyma-system/eventing-nats-0                                                 | 45 Mi    | 45 Mi   | 45 Mi    | 45 Mi    | 49 Mi    | -        | -        |
| kyma-system/eventing-nats-1                                                 | 81 Mi    | 81 Mi   | 81 Mi    | 83 Mi    | 85 Mi    | -        | -        |
| kyma-system/eventing-nats-2                                                 | 51 Mi    | 53 Mi   | 54 Mi    | 54 Mi    | 56 Mi    | -        | -        |
| kyma-system/istio-controller-manager-6b7f78d558-rpxf5                       | 60 Mi    | 60 Mi   | 63 Mi    | 78 Mi    | 72 Mi    | -        | -        |
| kyma-system/keda-admission-webhooks-855f88bf8-4c5kn                         | 12 Mi    | 12 Mi   | 12 Mi    | 13 Mi    | 13 Mi    | -        | -        |
| kyma-system/keda-manager-66d9959c47-mgjdf                                   | 83 Mi    | 83 Mi   | 83 Mi    | 83 Mi    | 83 Mi    | -        | -        |
| kyma-system/keda-operator-67b88dbf8b-pkgsg                                  | 50 Mi    | 344 Mi  | 690 Mi   | 768 Mi   | FAILING  | FAILING  | FAILING  |
| kyma-system/keda-operator-metrics-apiserver-6cbf7567fc-877pj                | 48 Mi    | 48 Mi   | 48 Mi    | 49 Mi    | 49 Mi    | -        | -        |
| kyma-system/kim-snatch-controller-manager-84cdb5df56-nwfp9                  | 46 Mi    | 46 Mi   | 46 Mi    | 47 Mi    | 47 Mi    | -        | -        |
| kyma-system/nats-manager-785744947f-w6rcb                                   | 90 Mi    | 91 Mi   | 86 Mi    | 88 Mi    | 88 Mi    | -        | -        |
| kyma-system/rma-kube-state-metrics-6879cd7599-q9254                         | 2756 Mi  | 2775 Mi | 2760 Mi  | 2786 Mi  | 2730 Mi  | -        | -        |
| kyma-system/rma-system-logs-agent-48hwt                                     | 31 Mi    | 31 Mi   | 31 Mi    | 32 Mi    | 31 Mi    | -        | -        |
| kyma-system/rma-system-logs-agent-4d4nr                                     | 31 Mi    | 31 Mi   | 31 Mi    | 32 Mi    | -        | -        | -        |
| kyma-system/rma-system-logs-agent-4khk5                                     | 46 Mi    | 46 Mi   | 47 Mi    | 48 Mi    | 47 Mi    | -        | -        |
| kyma-system/rma-system-logs-agent-9pwkx                                     | 36 Mi    | 36 Mi   | 36 Mi    | 36 Mi    | 36 Mi    | -        | -        |
| kyma-system/rma-system-logs-agent-c8vmd                                     | 32 Mi    | 32 Mi   | 33 Mi    | 33 Mi    | 33 Mi    | -        | -        |
| kyma-system/rma-system-logs-agent-hxw25                                     | 78 Mi    | 79 Mi   | 80 Mi    | 80 Mi    | 80 Mi    | -        | -        |
| kyma-system/rma-system-logs-agent-vvnlm                                     | 38 Mi    | 38 Mi   | 38 Mi    | 39 Mi    | 40 Mi    | -        | -        |
| kyma-system/rma-system-logs-agent-xzgr7                                     | 34 Mi    | 34 Mi   | 34 Mi    | 35 Mi    | 35 Mi    | -        | -        |
| kyma-system/rma-system-logs-collector-565d78d588-vxmsv                      | 64 Mi    | 64 Mi   | 64 Mi    | 64 Mi    | 64 Mi    | -        | -        |
| kyma-system/rma-victoria-metrics-agent-77ff8d586c-h5zzp                     | 1063 Mi  | 1118 Mi | 1140 Mi  | 1162 Mi  | 1164 Mi  | -        | -        |
| kyma-system/sap-btp-operator-controller-manager-74bd4cbdc6-7mk86            | 31 Mi    | 31 Mi   | 31 Mi    | 30 Mi    | 32 Mi    | -        | -        |
| kyma-system/serverless-ctrl-mngr-7d4bf94d45-7gt2f                           | 40 Mi    | 40 Mi   | 40 Mi    | 40 Mi    | 41 Mi    | -        | -        |
| kyma-system/serverless-operator-b5c797f8c-v7jtd                             | 60 Mi    | 60 Mi   | 60 Mi    | 61 Mi    | 65 Mi    | -        | -        |
| kyma-system/skr-webhook-5444569bf6-2l4hm                                    | -        | -       | -        | -        | 10 Mi    | -        | -        |
| kyma-system/skr-webhook-5444569bf6-5fmkm                                    | -        | -       | -        | 11 Mi    | -        | -        | -        |
| kyma-system/skr-webhook-5444569bf6-7dn9r                                    | -        | 22 Mi   | 22 Mi    | 22 Mi    | -        | -        | -        |
| kyma-system/skr-webhook-5444569bf6-b6g8w                                    | -        | -       | -        | -        | 8 Mi     | -        | -        |
| kyma-system/skr-webhook-5444569bf6-g48td                                    | 25 Mi    | 25 Mi   | -        | -        | -        | -        | -        |
| kyma-system/skr-webhook-5444569bf6-j7q75                                    | -        | -       | -        | 9 Mi     | 10 Mi    | -        | -        |
| kyma-system/skr-webhook-5444569bf6-sqrg8                                    | -        | -       | -        | -        | 7 Mi     | -        | -        |
| kyma-system/skr-webhook-5444569bf6-xlkn8                                    | -        | -       | -        | -        | 8 Mi     | -        | -        |
| kyma-system/telemetry-manager-674b746498-zzzt8                              | 61 Mi    | 64 Mi   | 65 Mi    | 66 Mi    | 66 Mi    | -        | -        |
| kyma-system/telemetry-metric-agent-4jlhk                                    | 318 Mi   | 318 Mi  | 310 Mi   | 313 Mi   | 316 Mi   | -        | -        |
| kyma-system/telemetry-metric-agent-d2rp9                                    | 355 Mi   | 355 Mi  | 356 Mi   | 358 Mi   | 355 Mi   | -        | -        |
| kyma-system/telemetry-metric-agent-dxrk5                                    | 293 Mi   | 295 Mi  | 295 Mi   | 291 Mi   | 295 Mi   | -        | -        |
| kyma-system/telemetry-metric-agent-nmzln                                    | 313 Mi   | 313 Mi  | 312 Mi   | 309 Mi   | 315 Mi   | -        | -        |
| kyma-system/telemetry-metric-agent-pf658                                    | 153 Mi   | 150 Mi  | 150 Mi   | 151 Mi   | 152 Mi   | -        | -        |
| kyma-system/telemetry-metric-agent-s8dvv                                    | 282 Mi   | 282 Mi  | 285 Mi   | 285 Mi   | 287 Mi   | -        | -        |
| kyma-system/telemetry-metric-agent-v5dk9                                    | 270 Mi   | 270 Mi  | 269 Mi   | 275 Mi   | 272 Mi   | -        | -        |
| kyma-system/telemetry-metric-agent-wzhgt                                    | 108 Mi   | 110 Mi  | 110 Mi   | 107 Mi   | -        | -        | -        |
| kyma-system/telemetry-metric-gateway-f478fb4c5-t2h2g                        | 230 Mi   | 232 Mi  | 226 Mi   | 227 Mi   | 235 Mi   | -        | -        |
| kyma-system/telemetry-metric-gateway-f478fb4c5-t9sr9                        | 125 Mi   | 120 Mi  | 121 Mi   | 124 Mi   | 123 Mi   | -        | -        |
| kyma-system/telemetry-self-monitor-5994d4c965-jq8np                         | 74 Mi    | 74 Mi   | 76 Mi    | 80 Mi    | 94 Mi    | -        | -        |
| kyma-system/warden-admission-767f55476-bzbfg                                | 50 Mi    | 50 Mi   | 50 Mi    | 49 Mi    | 51 Mi    | -        | -        |
| kyma-system/warden-operator-69bf7ff597-rv5bv                                | 236 Mi   | 236 Mi  | 236 Mi   | 238 Mi   | 238 Mi   | -        | -        |
| movies-rest/kwiatekus-movies-rest-7c7fc75756-ksz76                          | 253 Mi   | 253 Mi  | 254 Mi   | 258 Mi   | 259 Mi   | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2cd7c                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2n7v2                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-42xrp                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-4c9jp                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-4hdvt                     | -        | -       | -        | 7 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-4rw65                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-4zwx6                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-52txv                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-5bc5n                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-5cs88                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-5l87c                     | -        | -       | 6 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-5zzdl                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-646p2                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-66v6h                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-6bdlc                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-6fchp                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-6q4dv                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-6xpd7                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-6xq9l                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-7cbg9                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-7hz6n                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-7qjsf                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-8bvpk                     | -        | -       | 7 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-8fjvg                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-8mdrz                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-8qrqm                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-8twvt                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-96d4f                     | -        | -       | 7 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9jffx                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9jxfl                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9rqfv                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9tt6j                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-b2p6p                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-b4vzm                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-bd56q                     | -        | -       | -        | -        | 7 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-btjjs                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-btrcz                     | -        | -       | -        | -        | 7 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-btw78                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-c22zk                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-cbgtr                     | 8 Mi     | -       | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-ccn7p                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-cjkjq                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-cqfxv                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-cx5kw                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-dg572                     | -        | -       | 6 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-dhmkr                     | -        | -       | 7 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-dll8l                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-ds9vq                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-f28xd                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-fbdng                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-fm2qq                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-fpb5p                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-gjtgn                     | -        | -       | -        | -        | 7 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-h9gt2                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-hfb2c                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-hqjqb                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-hr9sg                     | -        | -       | 7 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-htncx                     | 8 Mi     | -       | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-hzxnx                     | 8 Mi     | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-j29qw                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-jjgp6                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-jjvxl                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-jzkbg                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-k77wt                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-kd6l2                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-kd6pd                     | -        | -       | -        | 7 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-kh86w                     | -        | -       | -        | -        | 7 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-l5ggl                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-l6qxw                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-ljd4p                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-lw58x                     | -        | 8 Mi    | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-lww6h                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-mbnnm                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-mrqfb                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-mwvjr                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-nq5lk                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-p859k                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-p95wq                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-pjmp2                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-pvb8m                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-q5bcz                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-q5gpl                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-q924k                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-qxj7v                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-r4zt4                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-r8qhp                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-rjmt8                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-rkrvl                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-rm4jl                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-s28mm                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-s5c8l                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-s5q4x                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-s989d                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-sln5f                     | -        | -       | 7 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-t9h6p                     | -        | 7 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-tj4l8                     | -        | -       | -        | -        | 6 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-tk2km                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-v9plq                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-vg22z                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-vjsvb                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-w6xfb                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-w8csx                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-wmlpf                     | -        | -       | -        | -        | 8 Mi     | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-wq2rh                     | -        | 8 Mi    | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-x5llc                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-x87jr                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-z5gm8                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-z5zwt                     | -        | 7 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-zdrph                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-zdxlc                     | -        | 7 Mi    | 7 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-zgd6j                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-zkhsp                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-zpjpg                     | -        | -       | 6 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-zwq88                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-producer-86547dcdfd-bfcdk                     | 11 Mi    | 8 Mi    | 13 Mi    | 11 Mi    | 12 Mi    | -        | -        |
| registry-proxy/registry-proxy-controller-5955cc9f9f-69x7f                   | 39 Mi    | 39 Mi   | 40 Mi    | 41 Mi    | 40 Mi    | -        | -        |
| registry-proxy/registry-proxy-operator-5f8dc4d9b7-z8mqb                     | 46 Mi    | FAILING | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| sap-transp-proxy-system/sap-transp-proxy-healthcheck-58966dd54b-znt8j       | 76 Mi    | 80 Mi   | 80 Mi    | 84 Mi    | 81 Mi    | -        | -        |
| sap-transp-proxy-system/sap-transp-proxy-manager-64586c989b-nffp5           | 84 Mi    | 84 Mi   | 86 Mi    | 92 Mi    | 88 Mi    | -        | -        |
| sap-transp-proxy-system/sap-transp-proxy-operator-76d5c8fd84-9zmcd          | 151 Mi   | 173 Mi  | 179 Mi   | 174 Mi   | 181 Mi   | -        | -        |
| ztis-agent-system/ztis-operator-controller-manager-c89c75d56-dvldl          | 67 Mi    | 67 Mi   | 67 Mi    | 70 Mi    | 71 Mi    | -        | -        |

