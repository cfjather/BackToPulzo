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
2. Instala las dependencias  

### Opcion No. 1: Instalacion Manual

3. Instala las dependencias manualmente
   - Para el frontend (Angular):  
     ```bash
     npm install
     ```  
   - Para el backend (Go):  
     Instala Go desde [https://go.dev/](https://go.dev/) o usa Chocolatey:  
     ```bash
     choco install golang
     ```  
3. Inicia el backend y el frontend manualmente:  
   - Backend:  
     ```bash
     go run main.go
     ```  
   - Frontend:  
     ```bash
     ng serve -o
     ```  

### Opcion No. 2: Docker
2. Asegúrate de tener Docker instalado.  
3. Desde la raíz del proyecto, ejecuta:

```bash
docker compose up --build
```

Puedes usarlo desde LocalHost, o via POSTMAN.

### Via LocalHost
4. Abre el servidor en tu navegador.
5. Para obtener un token, selecciona **Generador de Tokens** en la barra de navegación.  
6. Toca **Copiar Token** o cópialo manualmente.  
7. Ve a **Visor de Personajes**.  
8. Pega el token en el campo **Ingresa tu Token** y toca en **Obtener Personajes**. ¡Listo!  
9. Si deseas ver los tokens existentes, dirígete a **Administrador de Tokens**.  
10. Verás los tokens disponibles y sus usos restantes.  


### Via PostMan (Generacion de Tokens y Consulta de Personajes)
4. Configura el metodo en POST.
5. Toca 'ENVIAR'. Recibiras un JSON con el Token.
6. Configura ahora, el metodo en GET.
7. Dirigete a Headers.
8. En 'Key' utiliza 'Authorization'.
9. En Value utiliza el Token que generaste previamente.
10. Listo!

### Importante: El Administrador de Tokens solo esa disponible via HTML