package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	var opcion int

	//for {
	fmt.Println("\n\n\nBienvenido al programa para bloquear los puertos USB")
	fmt.Println("Seleccione una opcion:]")
	fmt.Println("1. Bloquear puertos USB")
	fmt.Println("2. Habilitar puertos USB")
	fmt.Println("3. salir")

	fmt.Scanln(&opcion)

	switch opcion {
	case 1:
		fmt.Println("Opcion para bloquear puertos USB")

		// Solicitar la contraseña del usuario y verificar si tiene permisos de administrador
		fmt.Println("Introduce tu contraseña para completar la accion: ")
		cmd := exec.Command("sudo", "-v")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Contraseña incorrecta")
			return
		}

		// Bloquear los puertos USB
		//blockCommand := "sudo sh -c 'echo ACTION==\"add\", SUBSYSTEMS==\"usb\", ATTRS{idVendor}==\"*\", ATTRS{idProduct}==\"*\", RUN+=\"/bin/sh -c 'echo 0 > /sys/bus/usb/devices/%k/authorized\"' >> /etc/udev/rules.d/99-block-usb.rules'"
		blockCommand := "sudo chmod 000 /media/"
		_, err := exec.Command("bash", "-c", blockCommand).Output()
		if err != nil {
			fmt.Println("Error")
			log.Fatal(err)
		}
		fmt.Println("****Puertos USB bloqueados")
	case 2:
		fmt.Println("Opcion para habilitar puertos USB")

		// Solicitar la contraseña del usuari3o y verificar si tiene permisos de administrador
		fmt.Println("Introduce tu contraseña para completar la accion: ")
		cmd := exec.Command("sudo", "-v")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Contraseña incorrecta")
			return
		}

		// Desbloquear los puertos USB
		//unblockCommand := "sudo sh -c 'echo ACTION==\"add\", SUBSYSTEMS==\"usb\", ATTRS{idVendor}==\"*\", ATTRS{idProduct}==\"*\", RUN+=\"/bin/sh -c 'echo 1 > /sys/bus/usb/devices/%k/authorized\"' >> /etc/udev/rules.d/99-unblock-usb.rules'"
		unblockCommand := "sudo chmod 777 /media/"
		_, err2 := exec.Command("bash", "-c", unblockCommand).Output()
		if err2 != nil {
			fmt.Println("Error")
			log.Fatal(err2)
		}
		fmt.Println("****Puertos USB habilitados")
	case 3:
		fmt.Println("Salir del programa")
		os.Exit(0)
	default:
		fmt.Println("Opcion no valida")
	}
	//}
}
