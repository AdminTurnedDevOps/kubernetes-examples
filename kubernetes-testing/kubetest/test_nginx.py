# how to run
# pytest -s .

import os

import pytest


def test_nginx(kube):
   """Test that a simple deployment runs as intended."""

   # Load and create a deployment
   deployment = kube.load_deployment('deployment.yaml')
   deployment.create()

   # Wait until the deployment is in the ready state and then
   # refresh its underlying object data
   deployment.wait_until_ready(timeout=10)
   deployment.refresh()

   # Get the pods from the deployment and check that we have
   # the right number of replicas
   pods = deployment.get_pods()
   assert len(pods) == 1

   # Get the pod, ensure that it is ready, then get the containers
   # for that pod.
   pod = pods[0]
   pod.wait_until_ready(timeout=10)

   containers = pod.get_containers()
   assert len(containers) == 1