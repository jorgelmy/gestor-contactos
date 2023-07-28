package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
}

func main() {
	fmt.Println("-------------------------------------------------------------------\n" +
		"-------------- Bienvenido a su gestor de contactos ----------------\n" +
		"-------------------------------------------------------------------")
	// Slice de contactos
	var contacts []Contact

	// Cargar contactos existentes desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos:", err)
	}

	// Crear instancia de fubio
	reader := bufio.NewReader(os.Stdin)

	for {
		// Mostrar opciones al usuario
		fmt.Print("\n==== Seleccione la opción de preferencia ====\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Buscar contacto por nombre\n",
			"4. Salir\n",
			"Elige una opción: ")
		// Leer la opción del usuario
		var option int
		_, optErr := fmt.Scanln(&option)
		if optErr != nil {
			fmt.Println("Error al leer la opción:", err)
			return
		}

		switch option {
		case 1:
			fmt.Println("Ingrese la información del contacto a agregar")
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Phone, _ = reader.ReadString('\n')
			// Agregar un contacto a Slice
			contacts = append(contacts, c)

			resError := saveContactsToFile(contacts)
			if resError != nil {
				fmt.Println("Error guardando")
				return
			}
		case 2:
			fmt.Println("====================================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Telefono: %s",
					index+1, contact.Name, contact.Phone)

				if index+1 != len(contacts) {
					fmt.Println("----------------------------------------------------")
				}
			}
			fmt.Println("====================================================")
		case 3:
			fmt.Println("Ingrese el nombre a buscar")
			var name string
			fmt.Scan(&name)
			fmt.Println("----------------------------------------------------")
			contacto, errLoad := loadContactByNameFromFile(name)
			if errLoad != nil || contacto.Name == "" {
				fmt.Println("Fallo la busqueda por nombre")
				fmt.Println("----------------------------------------------------")
				continue
			}
			fmt.Printf("Nombre: %sTelefono: %s",
				contacto.Name, contacto.Phone)
			fmt.Println("----------------------------------------------------")
		case 4:
			fmt.Println("¡Hasta la proxima!")
			return

		default:
			fmt.Println("Debe ingresar un digito entre 1 y 3")
		}
	}
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)

	if err != nil {
		return err
	}
	return nil
}

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

func loadContactByNameFromFile(name string) (Contact, error) {
	file, err := os.Open("contacts.json")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
	}
	defer file.Close()

	var cs []Contact

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cs)

	var c Contact
	name = strings.ToLower(name) + "\n"
	for _, con := range cs {
		if strings.ToLower(con.Name) == name {
			return con, nil
		}
	}

	if err != nil {
		return c, err
	}
	return c, nil
}
