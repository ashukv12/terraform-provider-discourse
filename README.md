# Terraform Discourse Provider

This terraform provider enables Create, Read, Update, Delete, and import operations for discourse users.

## Requirements

* [Go](https://golang.org/doc/install) >= 1.16 (To build the provider plugin) <br>
* [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x <br/>
* [Discourse](https://www.discourse.org/pricing)  account 

## Setup Discourse Account
 :heavy_exclamation_mark:  [IMPORTANT] : This provider can be successfully tested on any discourse account. <br><br>

1. Create a discourse account with your required subscription [Standard Plan/Business Account](https://www.discourse.org/pricing)<br>
2. Sign in to the [discourse account](https://www.discourse.org/)<br>
3. Go to `Dashboard`. Click on `API`. Create `New API Key`. For our purpose we need to create an API Key. <br>

This app will provide us with the API Key which will be needed to configure our provider and make request. <br>
 
## Initialise Discourse Provider in local machine 
1. Add the API Key and Username generted in Discourse App to respective fields in `main.tf` <br>
3. Run the following command :
 ```golang
cd terraform-provider-discourse
go mod init terraform-provider-discourse
```
4. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>

## Installation

For Linux:

1. Run the following command to create a vendor subdirectory which will comprise of  all provider dependencies. <br>
    ```
    ~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}
    ``` 
    Command: 
    ```bash
    export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"
    mkdir -p ~/.terraform.d/plugins/discourse.org/user/discourse/1.0/$OS_ARCH
    ```

2. Run `go build -o terraform-provider-discourse`. This will save the binary file in the main/root directory. <br>

3. Run this command to move this binary file to appropriate location. <br>
```
  mv terraform-provider-discourse ~/.terraform.d/plugins/discourse.org/user/discourse/1.0/$OS_ARCH
```    
 <p align="center">
 [OR]
 </p><br>

3. Manually move the file from current directory to destination directory.<br>
 


## Run the Terraform provider


### Basic Terraform Commands
1. `terraform init` - To initialize a working directory containing Terraform configuration files.
2. `terraform plan` - To create an execution plan. Displays the changes to be done.
3. `terraform apply` - To execute the actions proposed in a Terraform plan. Apply the chages.

#### Create User
1. Add the user `email`, `name`, and `active` in the respective field in `main.tf`
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that an account activation mail has been sent to the user.
5. Activate the account using the link provided in the mail, and you will see that a user has been successfully created.

#### Update the user
  User is allowed to update `name`, and `active`. Update the data, i.e., `name` or `active` of the user in the in the resource block of `main.tf` file and apply using `terraform apply`

#### Read the User Data
Add data and output blocks in the `main.tf` file and run `terraform plan` to read user data

#### Activate/Deactivate the user
Change the `active` field value from `false` to deactivate and `true` to activate and run `terraform apply`.

#### Delete the user
Delete the resource block of the particular user from `main.tf` file and run `terraform apply`.

#### Import a User Data
1. Write manually a resource configuration block for the User in `main.tf`, to which the imported object will be mapped.
2. Run the command `terraform import discourse_user.user1 [EMAIL_ID]`
3. Check for the attributes in the `.tfstate` file and fill them accordingly in resource block.


## Example Usage
```terraform
terraform {
  required_providers {
    discourse = {
      version = "1"
      source  = "discourse.org/user/discourse"
    }
  }
}

provider "discourse" {
  api_key = ""
  api_username = ""
  base_url = ""
}

resource "discourse_user" "user1" {
   email = "[EMAIL_ID]"
   name = "[NAME]"
   active = true
}

data "discourse_user" "user1" {
  email = "[EMAIL_ID]"
}

output "user1" {
  value = data.discourse_user.user1
}
```

## Undocumented APIs being used here:

#### Update a user by username

`https://{defaultHost}/u/{username}.json`

#### Activate the user

`http://{defaultHost}/admin/users/{USER_ID}/activate.json`

#### Get email id by using username 

`https://{defaultHost}/u/{username}/emails.json`


## Argument Reference:

* `email`       - (Required, String)  - The email address of the user.
* `name`           - (Optional, String)  - Name of the user in Discourse. 
* `active`         - (Optional, Boolean) - If set to false, the user will be deactivated.
