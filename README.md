# argocd-project-role-token-generator

This tool will help you generate tokens to authenticate against the argocd API.

## Build
 * Clone the project
 * Build the project :
````
argocd-project-role-token-generator $ go build
````

## Usage

````
Usage of ./token-generator:
  -lifetime duration
    	Lifetime of the token
  -project string
    	Argo CD project which the role belongs to
  -role string
    	Argo CD project role which to create a token for
````

Example :

````
argocd-project-role-token-generator $ ./token-generator -role github-c2c-ci -project odoo-integration
Argo CD secret key?
-> Enter the secret key of the argocd you want to authenticate to
````
