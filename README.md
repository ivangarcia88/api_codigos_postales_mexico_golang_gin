# Introducción

API desarrollada en Golang (Gin) para consultar colonia, municipio y estado sobre un código postal.

**Ejemplo del funcionamiento de la API** 

[https://golangcpapimexico-production.up.railway.app/code/45070](https://golangcpapimexico-production.up.railway.app/code/45070)

**Esta API emplea información provista por el sitio oficial:**

[https://www.correosdemexico.gob.mx/SSLServicios/ConsultaCP/CodigoPostal_Exportar.aspx](https://www.correosdemexico.gob.mx/SSLServicios/ConsultaCP/CodigoPostal_Exportar.aspx)

La información fue transformada a un archivo JSON mediante un script, el código para este proceso se encuentra en: 

[https://github.com/ivangarcia88/mexican_postal_code_reformat](https://github.com/ivangarcia88/mexican_postal_code_reformat)


Para iniciar el servicio, se recomienda emplear uvicorn

```plaintext
go run main.go
```
[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template/dTvvSf)
