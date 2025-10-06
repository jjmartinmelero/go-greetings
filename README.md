# 🎯 Go-Greetings Module

Un módulo Go que proporciona funciones para generar mensajes de saludo personalizados y aleatorios.

## 📋 Características

- ✨ Generación de saludos personalizados
- 🎲 Selección aleatoria de formatos de saludo
- 📝 Soporte para saludos individuales y múltiples
- ⚡ Manejo de errores robusto

## 🚀 Instalación

Para usar este módulo en tu proyecto, ejecuta:

```bash
go get github.com/jjmartinmelero/go-greetings
```

## 💡 Uso

### Importar el módulo

```go
import "github.com/jjmartinmelero/go-greetings"
```

### Ejemplos de uso

#### Saludo individual

```go
message, err := greetings.Hello("Juan")
if err != nil {
    log.Fatal(err)
}
fmt.Println(message)
```

#### Saludos múltiples

```go
names := []string{"Alice", "Bob", "Carol"}
messages, err := greetings.Hellos(names)
if err != nil {
    log.Fatal(err)
}
for name, message := range messages {
    fmt.Println(message)
}
```

## 📚 API Reference

### `Hello(name string) (string, error)`

Genera un saludo personalizado para un nombre específico.

- **Parámetros:**
  - `name string`: El nombre de la persona a saludar
- **Retorna:**
  - `string`: El mensaje de saludo
  - `error`: Error si el nombre está vacío

### `Hellos(names []string) (map[string]string, error)`

Genera saludos personalizados para una lista de nombres.

- **Parámetros:**
  - `names []string`: Lista de nombres
- **Retorna:**
  - `map[string]string`: Mapa de nombres a mensajes
  - `error`: Error si algún nombre está vacío

## 🧪 Tests

Para ejecutar los tests:

```bash
go test
```

## 📜 Requisitos

- Go 1.24.5 o superior

## 📝 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## 👨‍💻 Autor

- [jjmartinmelero](https://github.com/jjmartinmelero)

## 🤝 Contribuir

Las contribuciones son bienvenidas. Por favor, abre un issue para discutir los cambios que te gustaría hacer.

1. Fork el proyecto
2. Crea tu rama de características (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request