# api/v1/auth/login
```bash
curl --header "Content-Type: application/json"   \
--request POST \
--data '{"email":"","password":""}' \
http://localhost:9000/api/v1/auth/login
```

# api/v1/auth/register
```bash
curl --header "Content-Type: application/json"   \
--request POST \
--data '{"firstName":  "","lastName":  "","email":  "","password":  ""}' \
http://localhost:9000/api/v1/auth/register
```

# api/v1/auth/verify-account
```bash
curl --header "Content-Type: application/json" \
--request POST \
--data '{"email":  "","otp":  ""}' \
http://localhost:9000/api/v1/auth/verify-account
```

# api/v1/auth/forgot-password
```bash
curl --header "Content-Type: application/json" \
--request POST \
--data '{"email": "" }' \
http://localhost:9000/api/v1/auth/forgot-password
```

# api/v1/auth/reset-password
```bash
curl --header "Content-Type: application/json" \
--request POST \
--data '{"email":"","otp":"","password":""}' \
http://localhost:9000/api/v1/auth/reset-password
```


# api/v1/category/create
```bash
curl --header "Content-Type: application/json" \
--header 'Authorization: Bearer {token}' \
--request POST \
--data '{"name":""}' \
http://localhost:9000/api/v1/auth/category/create
```

# api/v1/category/list
```bash
curl --header "Content-Type: application/json" \
--header 'Authorization: Bearer {token}' \
--request POST \
http://localhost:9000/api/v1/category/list
```


# api/v1/project/create
```bash
curl --header "Content-Type: application/json" \
--header 'Authorization: Bearer {token}' \
--request POST \
--data '{"name":""}' \
http://localhost:9000/api/v1/project/create
```


# api/v1/project/list
```bash
curl --header "Content-Type: application/json" \
--header 'Authorization: Bearer {token}' \
--request POST \
http://localhost:9000/api/v1/project/list
```

# api/v1/project/report-list
```bash
curl --header "Content-Type: application/json" \
--header 'Authorization: Bearer {token}' \
--request POST \
http://localhost:9000/api/v1/project/report-list
```

# api/v1/expense/create
```bash
curl --header "Content-Type: application/json" \
--header 'Authorization: Bearer {token}' \
--request POST \
--data '{
"projectId":"",
"merchantName":"",
"amount":0,
"date":"",
"description":"",
"categoryId":"",
"includeVat":false,
"vat":0.0}' \
http://localhost:9000/api/v1/expense/create
```

# api/v1/expense/list
```bash
curl --header "Content-Type: application/json" \
--header 'Authorization: Bearer {token}' \
--request POST \
http://localhost:9000/api/v1/expense/list
```
