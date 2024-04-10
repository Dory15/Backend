
# Backend

Backend desarrollado en golang usando base de datos SQL Server


## API Reference

#### Registro

```http
  POST /signup
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `correo` | `string` | **Requerido**. Correo electrónico del usuario |
| `usuario` | `string` | **Requerido**. Usuario a utilizar |
| `contrasena` | `string` | **Requerido**. Password del usuario |
| `telefono` | `string` | **Requerido**. CTelefono del usuario |

Ejemplo
```json
{
    "correo": "dorian@dorian.com",
    "usuario": "dorian",
    "contrasena": "Dorian1$",
    "telefono": "5549204786"
}
```

Respuesta en caso de error
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `error` | `string` | Mensaje de error en caso de no haber podido registrar al usuario |

Ejemplo
```json
{
    "error": "El correo/telefono ya se encuentra registrado",
}
```


#### Inicio de sesión

```http
  POST /signin
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `usuarioocorreo` | `string` | **Requerido**. Correo electrónico o usuario de la cuenta a utilizar |
| `contrasena` | `string` | **Requerido**. Password del usuario |

Ejemplo
```json
{
    "usuarioocorreo": "dorian@dorian.com",
    "contrasena": "Dorian1$"
}
```

Respuesta 
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `error` | `string` | Mensaje de error en caso de no haber podido iniciar sesión |
| `token` | `string` | Token jwt en caso de que el inicio de sesión haya sido existoso |

Ejemplo
```json
{
    "error": "",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRvcmlhbkBkb3JpYW4uY29tIiwiZXhwIjoxNzEyOTQwNzc0fQ.tc3QS02igEj3Wz6OhAPnMNIsIxUM4hwOqKv_WCLayPw"
}
```

## Deployment

Para desplegar este proyecto de la forma más sencilla se utiliza docker compose con la imagen `dory15/backend_db:v1.0.0`, ya que la imagen contiene la base de datos Backend

```bash
  docker compose up -d
```

Si se desea hacer desde 0, cambiar el nombre de la imagen por `mcr.microsoft.com/mssql/server:2022-latest` en el docker compose, despues crear la base de datos Backend y luego cargar el contenido de los archivos de la carpeta database


## Environment Variables

Las variables de entorno del proyoecto son las siguientes:

`ENV` define el ambiente, en este caso prod

`DB_STRING` la cadena para conectarse a la base de datos

`PORT` el puerto por donde sale el servicio

`JWT_SECRET` la cadena para firmar los token

`JWT_EXPIRATION` y el tiempo en segundos de expiración del token