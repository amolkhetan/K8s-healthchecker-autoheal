#!/bin/bash

NAMESPACE="autoheal"
RELEASE="autoheal"

echo "=============================="
echo "✅ Validating Helm Installation"
echo "=============================="
helm version
helm list -n $NAMESPACE
helm status $RELEASE -n $NAMESPACE

echo
echo "=============================="
echo "✅ Checking Kubernetes Resources"
echo "=============================="
kubectl get all -n $NAMESPACE

echo
echo "=============================="
echo "✅ Checking Pod Status"
echo "=============================="
kubectl get pods -n $NAMESPACE | awk '{print $1, $3, $4, $5}'

# Optional: Describe pods not in Running state
NOT_RUNNING=$(kubectl get pods -n $NAMESPACE --field-selector=status.phase!=Running -o name)
if [ ! -z "$NOT_RUNNING" ]; then
  echo
  echo "⚠ Pods not running:"
  for pod in $NOT_RUNNING; do
    kubectl describe $pod -n $NAMESPACE
    kubectl logs $pod -n $NAMESPACE --tail=20
  done
else
  echo "All pods are running ✅"
fi

echo
echo "=============================="
echo "✅ Checking Services"
echo "=============================="
kubectl get svc -n $NAMESPACE

echo
echo "=============================="
echo "✅ Testing Internal Connectivity"
echo "=============================="
echo "Running a temporary pod to test service endpoints..."
kubectl run testpod --rm -i --tty --image=busybox -n $NAMESPACE -- sh -c "
echo 'Inside testpod:';
for svc in \$(kubectl get svc -n $NAMESPACE -o jsonpath='{.items[*].metadata.name}'); do
  PORT=\$(kubectl get svc \$svc -n $NAMESPACE -o jsonpath='{.spec.ports[0].port}');
  echo -n '\$svc:\t';
  wget -qO- http://\$svc:\$PORT || echo 'Failed';
done
"

echo
echo "=============================="
echo "✅ Helm Test Hooks (if any)"
echo "=============================="
helm test $RELEASE -n $NAMESPACE || echo "No test hooks defined or tests failed"

echo
echo "=============================="
echo "✅ Validation Completed"
echo "=============================="

