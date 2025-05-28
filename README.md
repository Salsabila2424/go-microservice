#Jalankan Service
   cd encryption-service || cd auth-service
   go run main.go

#Auth
curl -X POST http://localhost:8081/login 
     -H "Content-Type: application/json" 
     -d '{"username":"admin","password":"password"}'

#Encript
curl -X POST http://localhost:8082/encrypt \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer Token" \
  -d '{"text":"Hello, World!"}'


#Decript
curl -X POST http://localhost:8082/decrypt \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer Token" \
     -d '{"encrypted_text":"xq6DxWrczlvuWFv1xoyxplLV1PqmVcjl/OxBqR0="}'