# Set background image


## Descripcion

Establecer una imagen de fondo en la pantalla del dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setBackgroundImage",
        "value":{
            "imageUrl": "http://www.example.com/image.jpg"
        }
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| imageUrl | string | URL de la imagen |




## Response Message
```json
{
    ...,
    "data": {
        "error" : false,
        "msg" : ""
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| error | boolean | Error |
| msg | string | Mensaje de error o log interno|
