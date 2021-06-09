# Terraform Discourse Provider

This provider allows to perform Create ,Read ,Update, Delete, Deactivate, and Activate discourse users.

## Requirements

* [Go](https://golang.org/doc/install) 1.16 <br>
* [Terraform](https://www.terraform.io/downloads.html) 0.13.x <br/>
* [Discourse](https://www.discourse.org/pricing) Standard account 

## Setup Discourse Account
 :heavy_exclamation_mark:  [IMPORTANT] : This provider can be successfully tested on any discourse standard account. <br><br>

1. Create a discourse account with your required subscription (Standard Plan/Business Account). (https://www.discourse.org/pricing)<br>
2. Sign in to the discourse account (https://www.discourse.org/)<br>
3. Go to `Dashboard`. Click on `API`. For our purpose we need to create an API Key. <br>

This app will provide us with the API Key which will be needed to configure our provider and make request. <br>
 
## Initialise Discourse Provider in local machine 
1. Add the API Key and Username generted in Discourse App to respective fields in `main.tf` <br>
3. Run the following command :
 ```golang
go mod init terraform-provider-discourse
```
4. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>

## Installation
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

3. Run this command to move this binary file to appropriate location.
 ```
 mv terraform-provider-discourse ~/.terraform.d/plugins/discourse.org/user/discourse/1.0/$OS_ARCH
 ``` 
    Otherwise you can manually move the file from current directory to destination directory.<br>

    [OR]

    1. Download required binaries <br>
    2. move binary `~/.terraform.d/plugins/[architecture name]/`


## Run the Terraform provider

#### Create User
1. Add the user email in the respective field in `main.tf`
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that an account activation mail has been sent to the user.
5. Activate the account using the link provided in the mail, and you will see that a user has been successfully created.

#### Update the user
Update the data, i.e., `name` and `active` of the user in the `main.tf` file and apply using `terraform apply`

#### Read the User Data
Add data and output blocks in the `main.tf` file and run `terraform plan` to read user data

#### Activate/Deactivate the user
Change the active value User from false to `deactivate` and true to `activate` and run `terraform apply`.

#### Delete the user
Delete the resource block of the particular user from `main.tf` file and run `terraform apply`.

#### Import a User Data
1. Write manually a resource configuration block for the User in `main.tf`, to which the imported object will be mapped.
2. Run the command `terraform import discourse_user.user1 [EMAIL_ID]`
3. Check for the attributes in the `.tfstate` file and fill them accordingly in resource block.


### Testing the Provider
1. Navigate to the test file directory.
2. Run command `go test` . This command will give combined test result for the execution or errors if any failure occur.
3. If you want to see test result of each test function individually while running test in a single go, run command `go test -v`
4. To check test cover run `go test -cover`

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
