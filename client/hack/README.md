# Script User Guide

This README documents:
* What update-crd-codegen.sh goes
* When and how to use it

## update-crd-codegen.sh

This is the script to update CRD yaml file under /client/config/crd/ based on types.go file, and also the
generated deepcopy routines in /client/apis/volumepopulator/v1alpha1/zz_generated.deepcopy.go.

Make sure to run this script after making changes to /client/apis/volumepopulator/v1alpha1/types.go.

Follow these steps to update the CRD:

* Run ./hack/update-crd-codegen.sh from client directory, a new yaml file should have been created under 
  ./config/crd/
