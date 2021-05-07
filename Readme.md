## Go mastering course...
#
10 sections of the course

#


##

Section 10 
- You have to generate your .pems files with a openssl 

example: 
  > openssl rep x509 -nodes -days 365 -new rsa:256 -keyout key.pem -out cert.pem 

365 day is valid the news key and cert


- Also have to generate the hydra.db to conteins their passwords 
  > Hydra\hydradblayer\passwordvault\fillvault> go run .\fillvault.go 

## 
