package main

import (
	"log"

	"github.com/google/gousb"
)

const (
	vendorID  gousb.ID = 0x0416
	productID gousb.ID = 0x5011
)

func main() {
	ctx := gousb.NewContext()
	defer ctx.Close()

	dev, err := ctx.OpenDeviceWithVIDPID(vendorID, productID)
	if err != nil {
		log.Fatalf("No se pudo abrir el dispositivo: %v", err)
	}
	if dev == nil {
		log.Fatalf("Impresora POS58 no encontrada")
	}
	defer dev.Close()

	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalf("Error al obtener configuración: %v", err)
	}
	defer cfg.Close()

	intf, err := cfg.Interface(0, 0)
	if err != nil {
		log.Fatalf("Error al abrir interfaz: %v", err)
	}
	defer intf.Close()

	// Busca endpoint OUT (bit 7 == 0)
	var epOut *gousb.OutEndpoint
	for _, ep := range intf.Setting.Endpoints {
		if ep.Address&0x80 == 0 {
			epOut, err = intf.OutEndpoint(int(ep.Address))
			if err != nil {
				log.Fatalf("Error al abrir endpoint: %v", err)
			}
			break
		}
	}
	if epOut == nil {
		log.Fatalf("No se encontró endpoint OUT")
	}

	send := func(data []byte) {
		if _, err := epOut.Write(data); err != nil {
			log.Fatalf("Error al escribir: %v", err)
		}
	}

	// Espera breve para “despertar” la impresora

	send([]byte("\n"))
	send([]byte("\x1B\x64\x01"))
	// Texto
	send([]byte("Hola AMOR\n"))
	send([]byte("\x1b\x45\x01Negrita ON -- TE AMO\x1b\x45\x00\n"))

	// Avanza 5 líneas (deja espacio para cortar)
	send([]byte("\x1B\x64\x05"))

	// Corte de papel (si tu modelo lo soporta)
	send([]byte("\x1d\x56\x00"))
}
