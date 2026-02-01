1.
(sumulando link que recibe cliente/usuario potencial)
curl "http://localhost:8080/user-info?token=valid-token"

2.
endpoint POST para recepcionar la seleccion
curl -i -v -X POST http://localhost:8080/user-selection \
  -H "Content-Type: application/json" \
  -d '{
    "monto": "3900",
    "selection": "vacaciones",
    "plazo": "3 Meses",
    "periodicidad": "quincenal"
  }'

3.
endpoint POST para registro (ultima vista)
curl -i -v -X POST http://localhost:8080/create-account \
  -H "Content-Type: application/json" \
  -d '{
    "id": "francisco.marquez@gmail.com",
    "password": "123456"
  }'
