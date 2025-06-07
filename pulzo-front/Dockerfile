FROM node:20-alpine AS builder

WORKDIR /app

# Copia package.json y package-lock.json
COPY package*.json ./

# Instala TODAS las dependencias (incluyendo devDependencies)
RUN npm ci

# Copia el código fuente
COPY . .

# Build de la aplicación (Angular 19 usa production por defecto)
RUN npm run build

# Debug: verifica que se creó el build y su contenido
RUN ls -la dist/
RUN ls -la dist/pulzo-front/ || echo "No existe dist/pulzo-front"
RUN find dist/ -type f || echo "No hay archivos en dist/"

# Segunda etapa - NGINX
FROM nginx:alpine

# Copia los archivos buildados
COPY --from=builder /app/dist/pulzo-front/browser /usr/share/nginx/html

# Debug: verifica que se copiaron correctamente
RUN ls -la /usr/share/nginx/html

# Copia configuración de nginx
COPY default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]