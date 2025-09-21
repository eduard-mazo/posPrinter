# POS Printer (ESC/POS) — Python & Go

Ejemplos para imprimir en una **impresora térmica POS58** mediante comandos **ESC/POS**  
conectada por **USB**.  

Compatible con Windows, Linux y macOS (requiere `libusb`).

---

## 📂 Contenido

| Carpeta / Archivo | Descripción |
|------------------|-------------|
| `python/` | Código en Python (`pyusb`) para enviar texto y comandos ESC/POS. |
| `go/` | Código en Go (`gousb`) para imprimir desde Go. |
| `.gitignore` | Archivos y carpetas que no deben versionarse. |

---

## 🚀 Requisitos

### Python
- Python 3.8+
- Paquete [`pyusb`](https://pypi.org/project/pyusb/)
- `libusb` instalado en el sistema  
  - Windows: instalar driver con [Zadig](https://zadig.akeo.ie/) → **WinUSB** o **libusb-win32**

### Go
- Go 1.19+
- Paquete [`github.com/google/gousb`](https://pkg.go.dev/github.com/google/gousb)
- `libusb` en tu sistema (en Windows, usa Zadig para tu impresora).

---

## 🐍 Uso en Python

```bash
cd python
pip install pyusb
python prin.py
```

Ajusta los valores `VENDOR_ID` y `PRODUCT_ID` según tu dispositivo USB.

## 🐹 Uso en Go

```bash
cd go
go mod tidy
go run main.go
```
---

- Antes de ejecutar, asegúrate de:
  - Conectar la impresora.
  - Haber configurado el driver **libusb-win32** o **WinUSB** con [Zadig](https://zadig.akeo.ie/) (Windows).
  - En Linux/macOS: instalar **libusb-1.0-0-dev** (o equivalente).

## 🧩 Notas

- Si la primera letra no aparece, agrega un salto de línea (**\n**) o un pequeño retardo antes del primer texto.

- Para avanzar papel manualmente, envía varias líneas en blanco o el comando ESC d n (**\x1B\x64<n>**).

- El comando de corte (**\x1D\x56\x00**) solo funciona en modelos con cortador automático.