# Back to Pulzo - Prueba Desarrollador Backend.
Gestión de Tokens y Llamado de Personajes Rick and Morty

Este proyecto es una aplicación con Backend (GO), y Frontend (Angular) que permite:

- Generar y copiar tokens de autenticación.
- Visualizar tokens activos con sus usos restantes.
- Consultar personajes de la API pública de Rick and Morty usando un token válido.
- Manejar errores y mostrar mensajes claros al usuario.

---

## Tecnologías utilizadas

- Angular 19 (Standalone Components)
- GO 1.24.3 para gestión de tokens y puente hacia la API de Rick and Morty
- HTML5 y CSS3 para estilos (con CSS externo)
- TypeScript

---

## Funcionalidades principales

### Generación de Tokens

- Botón para generar un token (tanto en HTML como POSTMAN).
- Copiar token al portapapeles con mensaje de confirmación.

### Consulta de Personajes Rick and Morty

- Input para ingresar un token válido.
- Muestra información detallada de cada personaje (nombre, estado, ubicación, especie, origen, género e imagen).
- Indicación de usos restantes del token.
- Manejo de errores para tokens inválidos o vencidos.

### Visualización de Tokens Activos

- Tabla que lista tokens activos con sus usos restantes.
- Botón para actualizar la lista.

---

## Cómo ejecutar el proyecto

1. Clona el repositorio.
2. Instala las dependencias (npm install & choco install go (o descarga GO desde https://go.dev/)).
3. Inicia el Backend con GO (go run main.go) y el Frontend (ng serve -o).

Puedes usarlo desde LocalHost, o via POSTMAN.

### Via LocalHost
4. Abre el servidor LocalHost.
5. Para obtener un Token, selecciona 'Generador de Tokens' en la barra de navegacion.
6. Toca 'Copiar Token' o manualmente, copialo.
7. Dirigete a 'Visor de Personajes' en la barra de navegacion.
8. Pega el Token en el Campo 'Ingresa tu Token' y toca en 'Obtener Personajes'. Listo!
9. Si deseas ver los Tokens existentes, dirigete a 'Administrador de Tokens' en la barra de navegacion.
10. Podras ver los tokens disponibles y sus usos restantes.

### Via PostMan (Generacion de Tokens y Consulta de Personajes)
4. Configura el metodo en POST.
5. Toca 'ENVIAR'. Recibiras un JSON con el Token.
6. Configura ahora, el metodo en GET.
7. Dirigete a Headers.
8. En 'Key' utiliza 'Authorization'.
9. En Value utiliza el Token que generaste previamente.
10. Listo!

### Importante: El Administrador de Tokens solo esa disponible via HTML