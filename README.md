Project as an example to create an automated pipeline that performs Continous integration, Continous delivery and Continous deployment. The purpose of this project is to create a basic API hosted in a cloud kubernetes cluster(Azure Kubernetes Services) and to automate all the actions of CI/CD process.

This implementation has to follow a Trunk branch strategy. Every time a push is done into a feature branch, the continous integration flow runs automatically the following tasks:

1. Run unit tests.
2. Execute Linter.
3. Generate exec file.
4. Build a new version of Docker image.
5. Push the new version to ACR.

Once a Merge request to main is approved all the continous delivery and continous deployment tasks will be executed:

1. Execute user acceptance tests.
2. Login to AKS.
3. Perform the deployment with the image generated on previous step.
