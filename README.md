# KUBE INIT ENV

```bash
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="--disable=traefik" sh -
```

```bash
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

```bash
helm repo add traefik https://traefik.github.io/charts
```

```bash
helm repo update
```

```bash
kubectl apply -f traefik-pvc.yam
```

```bash
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
```


```bash
helm install traefik traefik/traefik --namespace kube-system --create-namespace -f traefik-value.yaml
```

```bash
kubectl apply -f ingressroute-config.yaml
```

```bash
cd go-service

kubectl apply -f service-a.yaml

kubectl apply -f ingressroute.yaml
```
