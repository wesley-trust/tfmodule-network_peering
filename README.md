# tfmodule-network_peering
Terraform module for deploying network peering (Local/Global) in Azure, with Terratest Unit/Integration testing via Go, and Regula (OPA) Policy as Code scanning in an Azure DevOps Pipeline
## [CI/CD Pipeline](https://dev.azure.com/wesleytrust/Terraform/_build?definitionId=71)
Select a stage to view in Azure DevOps. *Note that 'Skipped' stages in the last run, will show as "unknown" by design.*
| Pipeline |
| :------: |
|[![Build Status](https://dev.azure.com/wesleytrust/Terraform/_apis/build/status/Modules/Deployments/tfmodule-network_peering?repoName=wesley-trust%2Ftfmodule-network_peering&branchName=main)](https://dev.azure.com/wesleytrust/Terraform/_build/latest?definitionId=71&repoName=wesley-trust%2Ftfmodule-network_peering&branchName=main)|
### Testing Stages
| Unit Tests | Integration Tests |
|  :-------: | :---------------: |
|[![Build Status](https://dev.azure.com/wesleytrust/Terraform/_apis/build/status/Modules/Deployments/tfmodule-network_peering?repoName=wesley-trust%2Ftfmodule-network_peering&branchName=main&stageName=Unit)](https://dev.azure.com/wesleytrust/Terraform/_build/latest?definitionId=71&repoName=wesley-trust%2Ftfmodule-network_peering&branchName=main)|[![Build Status](https://dev.azure.com/wesleytrust/Terraform/_apis/build/status/Modules/Deployments/tfmodule-network_peering?repoName=wesley-trust%2Ftfmodule-network_peering&branchName=main&stageName=Integration)](https://dev.azure.com/wesleytrust/Terraform/_build/latest?definitionId=71&repoName=wesley-trust%2Ftfmodule-network_peering&branchName=main)|