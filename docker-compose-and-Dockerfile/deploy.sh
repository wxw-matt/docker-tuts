SERVICE='go-service'
OUTPUT=`kubectl get svc $SERVICE 2>&1`
if [[ $? -eq 0 ]]
then
  echo "Service $SERVICE has been deployed already."
else
  echo "Deploying ..."
  kubectl apply -f go-service-k8s.yml
  kubectl expose deployment/$SERVICE  --type="NodePort" --port 8080
fi

