# auth
For create a kubernetes user jwt token

## using token

> create jwt bare token

```
curl "http://fist.sealyun.svc.cluster.local:8080/token?user=fanux&group=sealyun&group=develop" 

eyJhbGciOiJSUzI1NiIsImtpZCI6IkNnYzRPVEV5TlRVM0VnWm5hWFJvZFdJIn0.eyJpc3MiOiJodHRwczovL2RleC5leGFtcGxlLmNvbTo4MDgwIiwic3ViIjoiQ2djNE9URXlOVFUzRWdabmFYUm9kV0kiLCJhdWQiOiJleGFtcGxlLWFwcCIsImV4cCI6MTU1MTA5NzkwNiwiaWF0IjoxNTUwNzM3OTA2LCJlbWFpbCI6ImZodGpvYkBob3RtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJncm91cHMiOlsiZGV2Il0sIm5hbWUiOiJmYW51eCJ9.ZqKn461UW0aGtyjyqu2Dc5tiUzC-6eYLag542d3AvklUdZuw8i9XwyaUg_f1OAj0ZsEcOybOe9_PeGMaUYzU0OvlKPY-q2zbQVC-m6u6sQw6ZXx8pi0W8k4wQSJnMaOLddCfurlYufmr8kScDBQlnKapSR0F9mJzvpKkHD-XNshQKWhX3n03g7OfFgb4RuhLjKDNQnoGn7DfBNntibHlF9sPo0jC5JjqTZaGvoGmiRE4PAXwxA-RJifsWDNf_jW8lrDiY4NSO_3O081cia4N1GKht51q9W3eaNMvFDD9hje7abDdZoz9KPi2vc3zvgH7cNv0ExVHKaA0-dwAZgTx4g
```

> use token with curl

```
# curl -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IkNnYzRPVEV5TlRVM0VnWm5hWFJvZFdJIn0.eyJpc3MiOiJodHRwczovL2RleC5leGFtcGxlLmNvbTo4MDgwIiwic3ViIjoiQ2djNE9URXlOVFUzRWdabmFYUm9kV0kiLCJhdWQiOiJleGFtcGxlLWFwcCIsImV4cCI6MTU1MTA5NzkwNiwiaWF0IjoxNTUwNzM3OTA2LCJlbWFpbCI6ImZodGpvYkBob3RtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJncm91cHMiOlsiZGV2Il0sIm5hbWUiOiJmYW51eCJ9.ZqKn461UW0aGtyjyqu2Dc5tiUzC-6eYLag542d3AvklUdZuw8i9XwyaUg_f1OAj0ZsEcOybOe9_PeGMaUYzU0OvlKPY-q2zbQVC-m6u6sQw6ZXx8pi0W8k4wQSJnMaOLddCfurlYufmr8kScDBQlnKapSR0F9mJzvpKkHD-XNshQKWhX3n03g7OfFgb4RuhLjKDNQnoGn7DfBNntibHlF9sPo0jC5JjqTZaGvoGmiRE4PAXwxA-RJifsWDNf_jW8lrDiY4NSO_3O081cia4N1GKht51q9W3eaNMvFDD9hje7abDdZoz9KPi2vc3zvgH7cNv0ExVHKaA0-dwAZgTx4g" \
-k https://172.31.12.61:6443/api/v1/namespaces/default/pods
{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {

  },
  "status": "Failure",
  "message": "pods is forbidden: User \"https://fist.sealyun.svc.cluster.local:8443#fanux\" cannot list resource \"pods\" in API group \"\" in the namespace \"default\"",
  "reason": "Forbidden",
  "details": {
    "kind": "pods"
  },
  "code": 403
}
```

> use in kubeconfig file

```
kubectl config set-credentials fanux --token=eyJhbGciOiJSUzI1NiIsImtpZCI6IkNnYzRPVEV5TlRVM0VnWm5hWFJvZFdJIn0.eyJpc3MiOiJodHRwczovL2RleC5leGFtcGxlLmNvbTo4MDgwIiwic3ViIjoiQ2djNE9URXlOVFUzRWdabmFYUm9kV0kiLCJhdWQiOiJleGFtcGxlLWFwcCIsImV4cCI6MTU1MTEwMDI5MywiaWF0IjoxNTUwNzQwMjkzLCJlbWFpbCI6ImZodGpvYkBob3RtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJncm91cHMiOlsiZGV2Il0sIm5hbWUiOiJmYW51eCJ9.OAK4oIYqJszm1EACYW2neXTo738RW9kXFOIN5bOT4Z2CeKAvYqyOVKCWZf04xX45jwT78mATR3uas2YvRooDXlvxaD3K43ls4KBSG-Ofp-ynqlcVTpD3sUDqyux2iieNv4N6IyCv11smrU0lIlkrQC6oyxzTGae1FrJVGc5rHNsIRZHp2WrQvw83uLn_elHgUfSlsOq0cPtVONaAQWMAMi2DX-y5GCNpn1CDvudGJihqsTciPx7bj0AOXyiOznWhV186Ybk-Rgqn8h0eBaQhFMyNpwVt6oIP5pvJQs0uoODeRv6P3I3-AjKyuCllh9KDtlCVvSP4WtMUTfHQN4BigQ

kubectl config set-context fanux --cluster=kubernetes --user=fanux

kubectl config use-context fanux

kubectl config set-cluster kubernetes --server=https://127.0.0.1:6443 --insecure-skip-tls-verify=true
```

```
[root@iZj6cegflzze2l7fpcqoerZ ~]# kubectl get pod
Error from server (Forbidden): pods is forbidden: User "https://fist.sealyun.svc.cluster.local:8443#fanux" cannot list resource "pods" in API group "" in the namespace "default"
```

> bind a role

```
[root@iZj6cegflzze2l7fpcqoerZ ~]# cat rolebind.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: read-secrets-global
subjects:
- kind: User
  name: "fanux" # Name is case sensitive
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: cluster-admin  # using admin role
  apiGroup: rbac.authorization.k8s.io
```
```
[root@iZj6cegflzze2l7fpcqoerZ ~]# kubectl  --kubeconfig /etc/kubernetes/admin.conf create  -f rolebind.yaml # 用管理员的kubeconfig
clusterrolebinding.rbac.authorization.k8s.io/read-secrets-global created

[root@iZj6cegflzze2l7fpcqoerZ ~]# kubectl get pod 
No resources found.
```


## test shell

```
## rbac visit k8s apiserver
curl -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IkNnYzRPVEV5TlRVM0VnWm5hWFJvZFdJIn0.eyJpc3MiOiJodHRwczovL2RleC5leGFtcGxlLmNvbTo4MDgwIiwic3ViIjoiQ2djNE9URXlOVFUzRWdabmFYUm9kV0kiLCJhdWQiOiJleGFtcGxlLWFwcCIsImV4cCI6MTU1MTA5NzkwNiwiaWF0IjoxNTUwNzM3OTA2LCJlbWFpbCI6ImZodGpvYkBob3RtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJncm91cHMiOlsiZGV2Il0sIm5hbWUiOiJmYW51eCJ9.ZqKn461UW0aGtyjyqu2Dc5tiUzC-6eYLag542d3AvklUdZuw8i9XwyaUg_f1OAj0ZsEcOybOe9_PeGMaUYzU0OvlKPY-q2zbQVC-m6u6sQw6ZXx8pi0W8k4wQSJnMaOLddCfurlYufmr8kScDBQlnKapSR0F9mJzvpKkHD-XNshQKWhX3n03g7OfFgb4RuhLjKDNQnoGn7DfBNntibHlF9sPo0jC5JjqTZaGvoGmiRE4PAXwxA-RJifsWDNf_jW8lrDiY4NSO_3O081cia4N1GKht51q9W3eaNMvFDD9hje7abDdZoz9KPi2vc3zvgH7cNv0ExVHKaA0-dwAZgTx4g" \
 -k https://172.31.12.61:6443/api/v1/namespaces/default/pods
## auth token
curl  http://fist.sealyun.svc.cluster.local:8080/token?user=fanux&group=dev&group=test
## auth well
curl https://fist.sealyun.svc.cluster.local:8443/.well-known/openid-configuration --cacert ca.pem
## auth keys
curl -v https://fist.sealyun.svc.cluster.local:8443/keys --cacert ca.pem
```
