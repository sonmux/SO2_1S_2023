package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/rjeczalik/notify"
)

func main() {
	// Abre el archivo de texto donde se escribirán los nombres de los archivos copiados
	file, err := os.OpenFile("/home/vboxuser/Desktop/so2/logUSB/log/registro2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Crea un canal para recibir eventos de cambio de archivo
	c := make(chan notify.EventInfo, 1)

	// Comienza a escuchar los eventos de cambio de archivo en la unidad USB
	var ruta = "/media/vboxuser/Ubuntu 20_04_4 LTS amd64"
	err = notify.Watch(ruta, c, notify.Create)
	if err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	// Escucha los eventos de cambio de archivo y escribe los nombres de los archivos copiados en el archivo de texto
	for {
		select {
		case ei := <-c:
			// Obtiene la ruta completa del archivo copiado
			fullPath := ei.Path()
			// Obtiene el nombre del archivo
			filename := filepath.Base(fullPath)
			// Escribe el nombre del archivo en el archivo de texto
			_, err := file.WriteString(fmt.Sprintf("archivo \"%s\" copiado en la ruta "+ruta+"\n", filename))
			if err != nil {
				log.Fatal(err)
			}
			// Imprime el nombre del archivo en la consola
			fmt.Printf("archivo \"%s\" copiado en la ruta "+ruta+"\n", filename)
		}
	}
}
