# 🚀 CLI GitHub API

CLI para interactuar con la **API de GitHub** de manera rápida y sencilla.

## 🔐 Configuración

El token de GitHub se puede cargar desde una **variable de entorno** o un archivo. Ejemplo usando variable de entorno:

```bash
export GH_TOKEN=ghp_tu_token_personal🧑🏻‍💻
```

## 💻 Instalación

1. Clona el repositorio:

```bash
git clone <URL-del-repositorio>
```

2. Accede al directorio del proyecto:

```bash
cd <nombre-del-directorio>
```

## 🚀 Ejecución

Para ejecutar el proyecto localmente:

```bash
go run main.go
```

## 🛠 Comandos disponibles

Actualmente, el CLI soporta los siguientes comandos:

| Comando | Descripción                           |
| ------- | ------------------------------------- |
| `1`     | Lista los miembros de la organización |
| `2`     | Lista repositorios de la organización |
| `3`     | Otorga permisos a colaborador         |
