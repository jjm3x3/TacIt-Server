$projectRoot=Split-Path $(pwd)

kubectl create secret generic cloudsql-instance-credentials --from-file=credentials.json=$projectRoot\config\tacit-db67b0097365.json
kubectl create secret generic cloudsql-db-credentials --from-literal=username=gorm --from-file=./password.txt
