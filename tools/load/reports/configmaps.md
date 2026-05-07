# ConfigMaps

## Summary

Pods with memory growth exceeding 200 Mi from baseline to 2000 resources (last measured column; pods that failed before 2000 res show delta at last valid measurement):

| pod (namespace/name)                                                        | baseline | 500 res | 1000 res | 2000 res | 3000 res | 4000 res | 5000 res | delta   |
| --------------------------------------------------------------------------- | -------- | ------- | -------- | -------- | -------- | -------- | -------- | ------- |
| istio-system/istiod-66b4585c8b-rnqzx                                        | 998 Mi   | 683 Mi  | 1057 Mi  | 1715 Mi  | -        | -        | -        | +717 Mi |
| kyma-system/istio-controller-manager-6b7f78d558-fmlph                       | 946 Mi   | 674 Mi  | 822 Mi   | 1454 Mi  | -        | -        | -        | +508 Mi |
| sap-transp-proxy-system/sap-transp-proxy-operator-76d5c8fd84-9zmcd          | 113 Mi   | 471 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  | +358 Mi |

Pods that have a valid baseline but start failing in further steps:

| pod (namespace/name)                                                        | baseline | 500 res | 1000 res | 2000 res | 3000 res | 4000 res | 5000 res |
| --------------------------------------------------------------------------- | -------- | ------- | -------- | -------- | -------- | -------- | -------- |
| sap-transp-proxy-system/sap-transp-proxy-operator-76d5c8fd84-9zmcd          | 113 Mi   | 471 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| cfapi-system/contour-contour-57fd8d96b8-mn452                               | 80 Mi    | 75 Mi   | 82 Mi    | 86 Mi    | -        | FAILING  | -        |

## Full report

| pod (namespace/name)                                                        | baseline | 500 res | 1000 res | 2000 res | 3000 res | 4000 res | 5000 res |
| --------------------------------------------------------------------------- | -------- | ------- | -------- | -------- | -------- | -------- | -------- |
| antek-system/fake-metrics-746484997-jncvm                                   | 10 Mi    | 10 Mi   | 10 Mi    | 10 Mi    | -        | -        | -        |
| antek-system/opensearch-external-scaler-fdbb7ffc-97zll                      | 8 Mi     | 8 Mi    | 8 Mi     | 8 Mi     | -        | -        | -        |
| cap-operator-system/cap-operator-controller-9c9bc97bb-5kpsc                 | 26 Mi    | 26 Mi   | 26 Mi    | 26 Mi    | -        | -        | -        |
| cap-operator-system/cap-operator-controller-manager-57cdf94999-dspsd        | FAILING  | FAILING | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| cap-operator-system/cap-operator-subscription-server-d997f7d8f-frq6k        | 9 Mi     | 9 Mi    | 9 Mi     | 9 Mi     | -        | -        | -        |
| cap-operator-system/cap-operator-webhook-786459c7ff-kl8pm                   | 11 Mi    | 11 Mi   | 11 Mi    | 11 Mi    | -        | -        | -        |
| cfapi-system/btp-service-broker-cc965b778-qjjpz                             | 15 Mi    | 15 Mi   | 15 Mi    | 15 Mi    | -        | -        | -        |
| cfapi-system/cfapi-operator-989bb945-4mn9s                                  | 166 Mi   | 166 Mi  | 173 Mi   | 198 Mi   | -        | -        | -        |
| cfapi-system/contour-contour-57fd8d96b8-mn452                               | 80 Mi    | 75 Mi   | 82 Mi    | 86 Mi    | -        | FAILING  | -        |
| cfapi-system/contour-envoy-2v5vq                                            | 32 Mi    | 32 Mi   | 32 Mi    | 32 Mi    | -        | -        | -        |
| cfapi-system/contour-envoy-7x65v                                            | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | -        | -        | -        |
| cfapi-system/contour-envoy-xbjgh                                            | 34 Mi    | 34 Mi   | 34 Mi    | 34 Mi    | -        | -        | -        |
| cfapi-system/contour-envoy-zdqns                                            | 35 Mi    | 35 Mi   | 35 Mi    | 35 Mi    | -        | -        | -        |
| demo-app/http-echo-677d479d69-s24pf                                         | 79 Mi    | 79 Mi   | 75 Mi    | 77 Mi    | -        | -        | -        |
| dev/api-postgresql-go-67f4797b5c-wppkk                                      | 72 Mi    | 72 Mi   | 68 Mi    | 73 Mi    | -        | -        | -        |
| dev/commerce-mock-75648c645b-lh26b                                          | 188 Mi   | 188 Mi  | 185 Mi   | 184 Mi   | -        | -        | -        |
| dev/fe-ui5-postgresql-55ff69fbd7-dhqgm                                      | 70 Mi    | 71 Mi   | 75 Mi    | 72 Mi    | -        | -        | -        |
| docker-registry/dockerregistry-fd5f6685c-tfxws                              | 80 Mi    | 80 Mi   | 80 Mi    | 80 Mi    | -        | -        | -        |
| docker-registry/dockerregistry-operator-587f69dbc4-rtz9z                    | 43 Mi    | 43 Mi   | 44 Mi    | 43 Mi    | -        | -        | -        |
| istio-system/istio-cni-node-2fll2                                           | 24 Mi    | 24 Mi   | 24 Mi    | 23 Mi    | -        | -        | -        |
| istio-system/istio-cni-node-gsf2b                                           | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | -        | -        | -        |
| istio-system/istio-cni-node-h26x8                                           | 30 Mi    | 30 Mi   | 30 Mi    | 29 Mi    | -        | -        | -        |
| istio-system/istio-cni-node-httjn                                           | 48 Mi    | 47 Mi   | 48 Mi    | 48 Mi    | -        | -        | -        |
| istio-system/istio-cni-node-p6qxd                                           | 31 Mi    | 31 Mi   | 29 Mi    | 31 Mi    | -        | -        | -        |
| istio-system/istio-cni-node-vzr4g                                           | 22 Mi    | 22 Mi   | 23 Mi    | 23 Mi    | -        | -        | -        |
| istio-system/istio-cni-node-wsspg                                           | 31 Mi    | 31 Mi   | 31 Mi    | 31 Mi    | -        | -        | -        |
| istio-system/istio-ingressgateway-76974d6f7f-65kbd                          | 99 Mi    | 99 Mi   | 100 Mi   | 99 Mi    | -        | -        | -        |
| istio-system/istio-ingressgateway-76974d6f7f-c5h5v                          | 72 Mi    | 72 Mi   | 68 Mi    | 72 Mi    | -        | -        | -        |
| istio-system/istio-ingressgateway-76974d6f7f-thkxr                          | 155 Mi   | 157 Mi  | 157 Mi   | 161 Mi   | -        | -        | -        |
| istio-system/istiod-66b4585c8b-njhj4                                        | 1199 Mi  | 975 Mi  | 1284 Mi  | 1284 Mi  | -        | -        | FAILING  |
| istio-system/istiod-66b4585c8b-rnqzx                                        | 998 Mi   | 683 Mi  | 1057 Mi  | 1715 Mi  | -        | -        | -        |
| korifi/korifi-api-deployment-6d8c9f784d-lr6kh                               | 17 Mi    | 17 Mi   | 17 Mi    | 17 Mi    | -        | -        | -        |
| korifi/korifi-controllers-controller-manager-5ddd86668-85tk6                | 970 Mi   | 970 Mi  | 970 Mi   | 969 Mi   | -        | -        | -        |
| kpack/kpack-controller-5c6866cf46-twcx6                                     | 46 Mi    | 46 Mi   | 46 Mi    | 46 Mi    | -        | -        | -        |
| kpack/kpack-webhook-54cf8f75bf-pk754                                        | 63 Mi    | 63 Mi   | 63 Mi    | 63 Mi    | -        | -        | -        |
| kube-system/apiserver-proxy-479wm                                           | 31 Mi    | 31 Mi   | 31 Mi    | 31 Mi    | -        | -        | -        |
| kube-system/apiserver-proxy-5lr8p                                           | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | -        | -        | -        |
| kube-system/apiserver-proxy-8kw9m                                           | 32 Mi    | 32 Mi   | 32 Mi    | 32 Mi    | -        | -        | -        |
| kube-system/apiserver-proxy-997dr                                           | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | -        | -        | -        |
| kube-system/apiserver-proxy-hf6qf                                           | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | -        | -        | -        |
| kube-system/apiserver-proxy-jpxjc                                           | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | -        | -        | -        |
| kube-system/apiserver-proxy-q688t                                           | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | -        | -        | -        |
| kube-system/blackbox-exporter-59694df6-tvrdm                                | 17 Mi    | 17 Mi   | 17 Mi    | 16 Mi    | -        | -        | -        |
| kube-system/blackbox-exporter-59694df6-vwn7t                                | 16 Mi    | 16 Mi   | 16 Mi    | 16 Mi    | -        | -        | -        |
| kube-system/calico-node-gw2x6                                               | 207 Mi   | 209 Mi  | 231 Mi   | 211 Mi   | -        | -        | -        |
| kube-system/calico-node-jg47c                                               | 214 Mi   | 214 Mi  | 230 Mi   | 214 Mi   | -        | -        | -        |
| kube-system/calico-node-kzvcd                                               | 207 Mi   | 207 Mi  | 209 Mi   | 226 Mi   | -        | -        | -        |
| kube-system/calico-node-lmvlz                                               | 215 Mi   | 215 Mi  | 216 Mi   | 221 Mi   | -        | -        | -        |
| kube-system/calico-node-ntsb5                                               | 198 Mi   | 219 Mi  | 216 Mi   | 202 Mi   | -        | -        | -        |
| kube-system/calico-node-vertical-autoscaler-57d6f4f8f-t6p6x                 | 12 Mi    | 12 Mi   | 12 Mi    | 13 Mi    | -        | -        | -        |
| kube-system/calico-node-vjp5x                                               | 200 Mi   | 201 Mi  | 204 Mi   | 198 Mi   | -        | -        | -        |
| kube-system/calico-node-xfhnq                                               | 197 Mi   | 198 Mi  | 197 Mi   | 197 Mi   | -        | -        | -        |
| kube-system/calico-typha-deploy-bd76c7b67-496gm                             | 113 Mi   | 114 Mi  | 109 Mi   | 111 Mi   | -        | -        | -        |
| kube-system/calico-typha-deploy-bd76c7b67-5k8hj                             | 97 Mi    | 97 Mi   | 96 Mi    | 96 Mi    | -        | -        | -        |
| kube-system/calico-typha-horizontal-autoscaler-66b75b5d95-rzspb             | 13 Mi    | 13 Mi   | 13 Mi    | 13 Mi    | -        | -        | -        |
| kube-system/calico-typha-vertical-autoscaler-54b5769b9-8fjp2                | 12 Mi    | 12 Mi   | 12 Mi    | 13 Mi    | -        | -        | -        |
| kube-system/cloud-node-manager-bw6nb                                        | 15 Mi    | 15 Mi   | 15 Mi    | 15 Mi    | -        | -        | -        |
| kube-system/cloud-node-manager-d4qkw                                        | 17 Mi    | 17 Mi   | 17 Mi    | 17 Mi    | -        | -        | -        |
| kube-system/cloud-node-manager-gmgxt                                        | 17 Mi    | 17 Mi   | 17 Mi    | 17 Mi    | -        | -        | -        |
| kube-system/cloud-node-manager-k25vj                                        | 36 Mi    | 36 Mi   | 36 Mi    | 36 Mi    | -        | -        | -        |
| kube-system/cloud-node-manager-krnf2                                        | 15 Mi    | 15 Mi   | 15 Mi    | 15 Mi    | -        | -        | -        |
| kube-system/cloud-node-manager-lrzq5                                        | 28 Mi    | 28 Mi   | 28 Mi    | 28 Mi    | -        | -        | -        |
| kube-system/cloud-node-manager-nzjvr                                        | 16 Mi    | 16 Mi   | 16 Mi    | 16 Mi    | -        | -        | -        |
| kube-system/coredns-57dd88d9c5-4mr4k                                        | 31 Mi    | 31 Mi   | 31 Mi    | 31 Mi    | -        | -        | -        |
| kube-system/coredns-57dd88d9c5-986ls                                        | 22 Mi    | 23 Mi   | 23 Mi    | 24 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-disk-62qkz                                      | 31 Mi    | 31 Mi   | 31 Mi    | 31 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-disk-7vzk5                                      | 37 Mi    | 37 Mi   | 37 Mi    | 37 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-disk-8vsm8                                      | 29 Mi    | 29 Mi   | 29 Mi    | 29 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-disk-gxksp                                      | 54 Mi    | 54 Mi   | 54 Mi    | 54 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-disk-hqrj7                                      | 32 Mi    | 32 Mi   | 32 Mi    | 32 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-disk-q7qns                                      | 31 Mi    | 31 Mi   | 31 Mi    | 31 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-disk-rpkkb                                      | 29 Mi    | 29 Mi   | 29 Mi    | 29 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-file-2zh57                                      | 27 Mi    | 27 Mi   | 27 Mi    | 27 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-file-49s9r                                      | 27 Mi    | 27 Mi   | 27 Mi    | 27 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-file-4cqsk                                      | 30 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-file-7tk8r                                      | 29 Mi    | 29 Mi   | 29 Mi    | 29 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-file-gfbhk                                      | 37 Mi    | 37 Mi   | 37 Mi    | 37 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-file-gx842                                      | 28 Mi    | 28 Mi   | 28 Mi    | 28 Mi    | -        | -        | -        |
| kube-system/csi-driver-node-file-j5nqx                                      | 40 Mi    | 40 Mi   | 40 Mi    | 40 Mi    | -        | -        | -        |
| kube-system/egress-filter-applier-bcc7b                                     | 15 Mi    | 15 Mi   | 15 Mi    | 14 Mi    | -        | -        | -        |
| kube-system/egress-filter-applier-ddz6s                                     | 8 Mi     | 8 Mi    | 8 Mi     | 8 Mi     | -        | -        | -        |
| kube-system/egress-filter-applier-fcwfv                                     | 11 Mi    | 11 Mi   | 11 Mi    | 11 Mi    | -        | -        | -        |
| kube-system/egress-filter-applier-k8bxb                                     | 9 Mi     | 9 Mi    | 9 Mi     | 9 Mi     | -        | -        | -        |
| kube-system/egress-filter-applier-mpcwc                                     | 15 Mi    | 15 Mi   | 15 Mi    | 15 Mi    | -        | -        | -        |
| kube-system/egress-filter-applier-q5zr5                                     | 12 Mi    | 12 Mi   | 12 Mi    | 12 Mi    | -        | -        | -        |
| kube-system/egress-filter-applier-t4g7n                                     | 12 Mi    | 12 Mi   | 12 Mi    | 12 Mi    | -        | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-69l5f                           | 31 Mi    | 31 Mi   | 30 Mi    | 28 Mi    | -        | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-cbd4s                           | 34 Mi    | 34 Mi   | 36 Mi    | 37 Mi    | -        | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-hrsbg                           | 27 Mi    | 27 Mi   | 30 Mi    | 30 Mi    | -        | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-lx7tt                           | 45 Mi    | 44 Mi   | 42 Mi    | 43 Mi    | -        | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-n25mv                           | 31 Mi    | 30 Mi   | 30 Mi    | 29 Mi    | -        | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-nkmzf                           | 30 Mi    | 30 Mi   | 30 Mi    | 29 Mi    | -        | -        | -        |
| kube-system/kube-proxy-cpu-worker-0-v1.34.6-qlhfv                           | 36 Mi    | 36 Mi   | 35 Mi    | 35 Mi    | -        | -        | -        |
| kube-system/metrics-server-bd6684798-d57vf                                  | 530 Mi   | 530 Mi  | 531 Mi   | 531 Mi   | -        | -        | -        |
| kube-system/metrics-server-bd6684798-psls2                                  | 336 Mi   | 336 Mi  | 337 Mi   | 337 Mi   | -        | -        | -        |
| kube-system/network-problem-detector-host-58tnx                             | 22 Mi    | 22 Mi   | 23 Mi    | 24 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-host-bz8jp                             | 21 Mi    | 22 Mi   | 23 Mi    | 23 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-host-gbwfs                             | 24 Mi    | 25 Mi   | 25 Mi    | 26 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-host-l6t7s                             | 21 Mi    | 22 Mi   | 24 Mi    | 24 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-host-p8728                             | 26 Mi    | 26 Mi   | 24 Mi    | 25 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-host-ts5jd                             | 20 Mi    | 20 Mi   | 20 Mi    | 20 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-host-xx48m                             | 24 Mi    | 23 Mi   | 24 Mi    | 23 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-pod-8svt5                              | 26 Mi    | 26 Mi   | 28 Mi    | 28 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-pod-c2wdm                              | 21 Mi    | 22 Mi   | 22 Mi    | 22 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-pod-c75rc                              | 24 Mi    | 24 Mi   | 23 Mi    | 24 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-pod-f4plc                              | 20 Mi    | 20 Mi   | 20 Mi    | 21 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-pod-hcxsv                              | 23 Mi    | 24 Mi   | 24 Mi    | 25 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-pod-hhbqr                              | 21 Mi    | 22 Mi   | 22 Mi    | 21 Mi    | -        | -        | -        |
| kube-system/network-problem-detector-pod-p44l8                              | 29 Mi    | 30 Mi   | 30 Mi    | 30 Mi    | -        | -        | -        |
| kube-system/node-exporter-cmgn6                                             | 16 Mi    | 16 Mi   | 17 Mi    | 17 Mi    | -        | -        | -        |
| kube-system/node-exporter-hbqp6                                             | 23 Mi    | 23 Mi   | 23 Mi    | 23 Mi    | -        | -        | -        |
| kube-system/node-exporter-lgpn7                                             | 14 Mi    | 14 Mi   | 14 Mi    | 14 Mi    | -        | -        | -        |
| kube-system/node-exporter-lzz27                                             | 15 Mi    | 15 Mi   | 15 Mi    | 15 Mi    | -        | -        | -        |
| kube-system/node-exporter-mmqph                                             | 16 Mi    | 16 Mi   | 17 Mi    | 17 Mi    | -        | -        | -        |
| kube-system/node-exporter-qsdhl                                             | 10 Mi    | 10 Mi   | 10 Mi    | 10 Mi    | -        | -        | -        |
| kube-system/node-exporter-sc76n                                             | 18 Mi    | 18 Mi   | 18 Mi    | 18 Mi    | -        | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-6khd2                               | 22 Mi    | 22 Mi   | 22 Mi    | 22 Mi    | -        | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-6sqvw                               | 24 Mi    | 25 Mi   | 25 Mi    | 25 Mi    | -        | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-8mblj                               | 22 Mi    | 22 Mi   | 22 Mi    | 23 Mi    | -        | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-klplf                               | 29 Mi    | 29 Mi   | 30 Mi    | 30 Mi    | -        | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-knndk                               | 25 Mi    | 25 Mi   | 25 Mi    | 28 Mi    | -        | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-nx6ht                               | 22 Mi    | 23 Mi   | 23 Mi    | 23 Mi    | -        | -        | -        |
| kube-system/node-local-dns-cpu-worker-0-r2rzc                               | 21 Mi    | 21 Mi   | 21 Mi    | 19 Mi    | -        | -        | -        |
| kube-system/node-problem-detector-74zcq                                     | 28 Mi    | 28 Mi   | 30 Mi    | 35 Mi    | -        | -        | -        |
| kube-system/node-problem-detector-cdqdq                                     | 31 Mi    | 29 Mi   | 28 Mi    | 36 Mi    | -        | -        | -        |
| kube-system/node-problem-detector-cgxjl                                     | 25 Mi    | 38 Mi   | 25 Mi    | 38 Mi    | -        | -        | -        |
| kube-system/node-problem-detector-jk7dp                                     | 26 Mi    | 24 Mi   | 28 Mi    | 28 Mi    | -        | -        | -        |
| kube-system/node-problem-detector-ls9d9                                     | 28 Mi    | 28 Mi   | 28 Mi    | 28 Mi    | -        | -        | -        |
| kube-system/node-problem-detector-phn5t                                     | 34 Mi    | 43 Mi   | 35 Mi    | 48 Mi    | -        | -        | -        |
| kube-system/node-problem-detector-w5v2s                                     | 34 Mi    | 22 Mi   | 21 Mi    | 35 Mi    | -        | -        | -        |
| kube-system/vpn-shoot-0                                                     | 16 Mi    | 16 Mi   | 17 Mi    | 16 Mi    | -        | -        | -        |
| kube-system/vpn-shoot-1                                                     | 13 Mi    | 13 Mi   | 13 Mi    | 13 Mi    | -        | -        | -        |
| kwok-system/kwok-controller-d5dbcc9d9-6g2n4                                 | 22 Mi    | 27 Mi   | 27 Mi    | 28 Mi    | -        | -        | -        |
| kyma-system/api-gateway-controller-manager-66d9d4d67d-f867w                 | 1190 Mi  | 1056 Mi | 1335 Mi  | 1335 Mi  | -        | -        | -        |
| kyma-system/application-connector-controller-manager-66f7775d77-wttlq       | FAILING  | FAILING | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| kyma-system/btp-manager-controller-manager-7c6d47d7b6-dgbld                 | 31 Mi    | 28 Mi   | 28 Mi    | 31 Mi    | -        | -        | -        |
| kyma-system/central-application-connectivity-validator-87c7b7947-m49hk      | 95 Mi    | 95 Mi   | 95 Mi    | 92 Mi    | -        | -        | -        |
| kyma-system/central-application-connectivity-validator-87c7b7947-rrc7t      | 79 Mi    | 80 Mi   | 80 Mi    | 76 Mi    | -        | -        | -        |
| kyma-system/central-application-gateway-7b695b9b57-jc94f                    | 79 Mi    | 79 Mi   | 75 Mi    | 79 Mi    | -        | -        | -        |
| kyma-system/central-application-gateway-7b695b9b57-tws2l                    | 80 Mi    | 80 Mi   | 81 Mi    | 80 Mi    | -        | -        | -        |
| kyma-system/compass-runtime-agent-567c6fbdb7-xjfv4                          | 94 Mi    | 94 Mi   | 94 Mi    | 91 Mi    | -        | -        | -        |
| kyma-system/connectivity-proxy-0                                            | 206 Mi   | 206 Mi  | 205 Mi   | 209 Mi   | -        | -        | -        |
| kyma-system/connectivity-proxy-operator-56c68f65c5-bpdj7                    | 390 Mi   | 391 Mi  | 391 Mi   | 391 Mi   | -        | -        | -        |
| kyma-system/connectivity-proxy-region-configurations-controller-b58ff5wvgx2 | 235 Mi   | 235 Mi  | 236 Mi   | 231 Mi   | -        | -        | -        |
| kyma-system/connectivity-proxy-restart-watcher-5656fc5865-gw495             | 359 Mi   | 359 Mi  | 361 Mi   | 361 Mi   | -        | -        | -        |
| kyma-system/connectivity-proxy-sm-operator-59dff4d5df-2rn9b                 | 220 Mi   | 220 Mi  | 220 Mi   | 220 Mi   | -        | -        | -        |
| kyma-system/eventing-nats-0                                                 | 44 Mi    | 44 Mi   | 45 Mi    | 44 Mi    | -        | -        | -        |
| kyma-system/eventing-nats-1                                                 | 80 Mi    | 82 Mi   | 82 Mi    | 82 Mi    | -        | -        | -        |
| kyma-system/eventing-nats-2                                                 | 54 Mi    | 54 Mi   | 51 Mi    | 53 Mi    | -        | -        | -        |
| kyma-system/istio-controller-manager-6b7f78d558-fmlph                       | 946 Mi   | 674 Mi  | 822 Mi   | 1454 Mi  | -        | -        | -        |
| kyma-system/keda-admission-webhooks-855f88bf8-4c5kn                         | 12 Mi    | 12 Mi   | 12 Mi    | 12 Mi    | -        | -        | -        |
| kyma-system/keda-manager-66d9959c47-mgjdf                                   | 69 Mi    | 69 Mi   | 69 Mi    | 69 Mi    | -        | -        | -        |
| kyma-system/keda-operator-67b88dbf8b-pkgsg                                  | 50 Mi    | 50 Mi   | 50 Mi    | 50 Mi    | -        | -        | -        |
| kyma-system/keda-operator-metrics-apiserver-6cbf7567fc-877pj                | 48 Mi    | 48 Mi   | 48 Mi    | 48 Mi    | -        | -        | -        |
| kyma-system/kim-snatch-controller-manager-84cdb5df56-nwfp9                  | 21 Mi    | 20 Mi   | 20 Mi    | 21 Mi    | -        | -        | -        |
| kyma-system/nats-manager-785744947f-w6rcb                                   | 91 Mi    | 91 Mi   | 87 Mi    | 87 Mi    | -        | -        | -        |
| kyma-system/rma-kube-state-metrics-6879cd7599-q9254                         | 2899 Mi  | 2941 Mi | 2951 Mi  | 2909 Mi  | -        | -        | -        |
| kyma-system/rma-system-logs-agent-48hwt                                     | 31 Mi    | 31 Mi   | 30 Mi    | 31 Mi    | -        | -        | -        |
| kyma-system/rma-system-logs-agent-4khk5                                     | 31 Mi    | 32 Mi   | 32 Mi    | 32 Mi    | -        | -        | -        |
| kyma-system/rma-system-logs-agent-9pwkx                                     | 36 Mi    | 36 Mi   | 35 Mi    | 37 Mi    | -        | -        | -        |
| kyma-system/rma-system-logs-agent-c8vmd                                     | 31 Mi    | 31 Mi   | 32 Mi    | 33 Mi    | -        | -        | -        |
| kyma-system/rma-system-logs-agent-hxw25                                     | 35 Mi    | 35 Mi   | 35 Mi    | 35 Mi    | -        | -        | -        |
| kyma-system/rma-system-logs-agent-vvnlm                                     | 37 Mi    | 38 Mi   | 38 Mi    | 39 Mi    | -        | -        | -        |
| kyma-system/rma-system-logs-agent-xzgr7                                     | 40 Mi    | 39 Mi   | 39 Mi    | 40 Mi    | -        | -        | -        |
| kyma-system/rma-system-logs-collector-565d78d588-vxmsv                      | 63 Mi    | 63 Mi   | 63 Mi    | 63 Mi    | -        | -        | -        |
| kyma-system/rma-victoria-metrics-agent-77ff8d586c-h5zzp                     | 1123 Mi  | 1147 Mi | 1147 Mi  | 1157 Mi  | -        | -        | -        |
| kyma-system/sap-btp-operator-controller-manager-74bd4cbdc6-7mk86            | 41 Mi    | 42 Mi   | 44 Mi    | 44 Mi    | -        | -        | -        |
| kyma-system/serverless-ctrl-mngr-7d4bf94d45-7gt2f                           | 24 Mi    | 24 Mi   | 24 Mi    | 24 Mi    | -        | -        | -        |
| kyma-system/serverless-operator-b5c797f8c-v7jtd                             | 35 Mi    | 35 Mi   | 35 Mi    | 38 Mi    | -        | -        | -        |
| kyma-system/skr-webhook-5444569bf6-ckp5k                                    | 8 Mi     | 8 Mi    | -        | -        | -        | -        | -        |
| kyma-system/skr-webhook-5444569bf6-dbt6g                                    | -        | -       | 4 Mi     | 9 Mi     | -        | -        | -        |
| kyma-system/skr-webhook-5444569bf6-rmz9v                                    | -        | 4 Mi    | 9 Mi     | -        | -        | -        | -        |
| kyma-system/skr-webhook-5444569bf6-vfxc9                                    | -        | -       | 10 Mi    | -        | -        | -        | -        |
| kyma-system/telemetry-manager-674b746498-zzzt8                              | 48 Mi    | 43 Mi   | 43 Mi    | 45 Mi    | -        | -        | -        |
| kyma-system/telemetry-metric-agent-4jlhk                                    | 321 Mi   | 316 Mi  | 314 Mi   | 317 Mi   | -        | -        | -        |
| kyma-system/telemetry-metric-agent-d2rp9                                    | 303 Mi   | 303 Mi  | 301 Mi   | 294 Mi   | -        | -        | -        |
| kyma-system/telemetry-metric-agent-dxrk5                                    | 285 Mi   | 285 Mi  | 287 Mi   | 287 Mi   | -        | -        | -        |
| kyma-system/telemetry-metric-agent-nmzln                                    | 312 Mi   | 312 Mi  | 308 Mi   | 310 Mi   | -        | -        | -        |
| kyma-system/telemetry-metric-agent-pf658                                    | 297 Mi   | 297 Mi  | 297 Mi   | 302 Mi   | -        | -        | -        |
| kyma-system/telemetry-metric-agent-s8dvv                                    | 280 Mi   | 281 Mi  | 281 Mi   | 281 Mi   | -        | -        | -        |
| kyma-system/telemetry-metric-agent-v5dk9                                    | 275 Mi   | 275 Mi  | 270 Mi   | 271 Mi   | -        | -        | -        |
| kyma-system/telemetry-metric-gateway-f478fb4c5-t2h2g                        | 217 Mi   | 216 Mi  | 209 Mi   | 212 Mi   | -        | -        | -        |
| kyma-system/telemetry-metric-gateway-f478fb4c5-t9sr9                        | 300 Mi   | 302 Mi  | 303 Mi   | 304 Mi   | -        | -        | -        |
| kyma-system/telemetry-self-monitor-5994d4c965-jq8np                         | 71 Mi    | 67 Mi   | 93 Mi    | 69 Mi    | -        | -        | -        |
| kyma-system/warden-admission-767f55476-bzbfg                                | 30 Mi    | 31 Mi   | 33 Mi    | 34 Mi    | -        | -        | -        |
| kyma-system/warden-operator-69bf7ff597-rv5bv                                | 222 Mi   | 225 Mi  | 226 Mi   | 226 Mi   | -        | -        | -        |
| movies-rest/kwiatekus-movies-rest-58f4876589-wngbj                          | 239 Mi   | 239 Mi  | 243 Mi   | 239 Mi   | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2cflj                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2ftkz                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2jclj                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2jdh7                     | -        | -       | -        | 7 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2s5p2                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-2t8s9                     | -        | 7 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-45qln                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-4jq4c                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-4vk5n                     | -        | -       | -        | 7 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-4vn4m                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-5668w                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-5wnj2                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-6sqwx                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-6vqkn                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-72s96                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-76w7f                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-7925c                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-7h5f6                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-7wz78                     | -        | -       | -        | 7 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-8rwrz                     | -        | -       | -        | 7 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-8zz2n                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9dbnc                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9j8t6                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9jdmh                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9rnlx                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-9rvqf                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-bm67p                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-ctblt                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-d2lv6                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-dqsz5                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-dt9j4                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-dwkjw                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-f4dz4                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-fct86                     | FAILING  | 6 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-ffz6x                     | -        | -       | 8 Mi     | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-fgh68                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-fr7b8                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-fz9t8                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-g5jzh                     | -        | -       | -        | 7 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-g62w9                     | -        | -       | -        | 7 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-gpwnt                     | 8 Mi     | -       | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-gxk4b                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-h4fxl                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-hh4pt                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-hw59w                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-j482w                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-k9kp5                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-kkvbr                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-km9vt                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-kpk9c                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-l2lg9                     | -        | -       | -        | 7 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-ldnjl                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-lffrg                     | FAILING  | 7 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-lh7ks                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-lhtvf                     | -        | -       | 8 Mi     | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-lv5nk                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-mbssj                     | 8 Mi     | -       | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-mffvf                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-nw5vd                     | -        | -       | 6 Mi     | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-nz7f9                     | 8 Mi     | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-nz7zt                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-p7nxw                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-pm5kg                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-psbp7                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-pwdjd                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-qbtkx                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-qmg6t                     | 6 Mi     | -       | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-qnc72                     | -        | 4 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-qqtb4                     | -        | 6 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-r5jtv                     | -        | -       | 7 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-r5m4k                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-r89jj                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-rmc5v                     | -        | 6 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-sp85s                     | -        | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-sp8nw                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-tt7gn                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-txx5f                     | 8 Mi     | 8 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-v5dqc                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-w568b                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-w6djj                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-wgbgz                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-wn966                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-wx8mk                     | -        | 1 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-x45zn                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-x4fg7                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-x8d4q                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-xfxrt                     | -        | -       | -        | 6 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-xr4j7                     | -        | 7 Mi    | -        | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-xv8p5                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-zpjjh                     | -        | -       | -        | 8 Mi     | -        | -        | -        |
| nats-jetstream-demo/jetstream-consumer-54bb874699-zpx58                     | -        | -       | 8 Mi     | -        | -        | -        | -        |
| nats-jetstream-demo/jetstream-producer-86547dcdfd-bfcdk                     | 6 Mi     | 1 Mi    | 1 Mi     | 8 Mi     | -        | -        | -        |
| registry-proxy/registry-proxy-controller-5955cc9f9f-g82cd                   | 37 Mi    | 37 Mi   | 36 Mi    | 36 Mi    | -        | -        | -        |
| registry-proxy/registry-proxy-operator-5f8dc4d9b7-z8mqb                     | 33 Mi    | 33 Mi   | 33 Mi    | 33 Mi    | -        | -        | -        |
| sap-transp-proxy-system/sap-transp-proxy-healthcheck-58966dd54b-znt8j       | 83 Mi    | 83 Mi   | 86 Mi    | 86 Mi    | -        | -        | -        |
| sap-transp-proxy-system/sap-transp-proxy-manager-64586c989b-nffp5           | 86 Mi    | 86 Mi   | 84 Mi    | 88 Mi    | -        | -        | -        |
| sap-transp-proxy-system/sap-transp-proxy-operator-76d5c8fd84-9zmcd          | 113 Mi   | 471 Mi  | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |
| ztis-agent-system/ztis-operator-controller-manager-c89c75d56-dvldl          | FAILING  | FAILING | FAILING  | FAILING  | FAILING  | FAILING  | FAILING  |

