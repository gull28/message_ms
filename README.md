# Microservice for user validation codes

## Running the service

- Clone the repository
- Run `go get` for installing go packages

### Caveats
If you wish to use email messaging only, the service is ready to go. Now, if you also want SMS handling, you will have to add a little bit of your own code based on the SMS provider you use. 
In most cases, this is very simple and all the information is provided by each specific SMS API provider. Simply edit the `sms.go` file with the specific code from the provider.

If you want to make a custom email with your own variables, use the `/delivery/templates` folder to store your html template to use in your email



