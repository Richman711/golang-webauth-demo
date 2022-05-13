# golang-webauth-demo
Demo of various Web authentication methods using Golang

```
# Tags are ran after a commit and push
# Adding Tags with the following command
git tag v0.0.0
git push --tags

# Tags are added after a commit to provide the ability 
# go to a version of code during a spacific commit
```

# Tags

## v0.2.0
JSON Marshal example

## v0.3.0
JSON Unmarshal example

## v0.4.0
Setup server for encodeing and decoding JSON

## v0.5.0
Encoding JSON
```
# After running go run main.go
# you need to curl the encode listener
curl localhost:8080/encode
```

## v0.6.0
Decode JSON
```
# Enter the following curl command to
# a seperate terminal after runing 
# go run main.go
# See the log output in the first terminal used
# to run main.go
#
# You can use https://curlbuilder.com/ to build 
# custom curl commands
curl -XGET -H "Content-type: application/json" -d '{"First":"Jonny"}' 'localhost:8080/decode'
```

## v0.7.0
HTTP Basic Authentication
  * Send username & password with every request
  * https://pkg.go.dev/net/http#Request.BasicAuth

## v0.8.0
Storing passwords
 * Always Always hash the password using a secure hashing algorithm. Never store raw passwords in an unhashed form.
  * bcrypt - Password hashing algorithm
    * Salting is achieved by 
    * https://pkg.go.dev/golang.org/x/crypto/bcrypt
  * scrypt - Password hashing
    * Includes a builtin salting method
    * https://pkg.go.dev/golang.org/x/crypto@v0.0.0-20220511200225-c6db032c6c88/scrypt
