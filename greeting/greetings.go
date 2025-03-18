package greeting

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	//check if name is empty
	if name == "" {
		return "", errors.New("empty name")
	}

	messageFormat := messageFormater()

	message := fmt.Sprintf(messageFormat, name)
	return message, nil
}

func messageFormater() string {

	formats := []string{
		" Hello, %v Welcome!",
		" Namaste, %v Swagat hai!",
		" namaskara, %v Suswagata!",
		" Hola, %v Bienvenido!",
		" Bonjour, %v Bienvenue!",
		" Ciao, %v Benvenuto!",
		" Ola, %v Bem-vindo!",
		" Guten Tag, %v Willkommen!",
		" Saluton, %v Bonvenon!",
		" Hej, %v Välkommen!",
		" Aloha, %v Welina!",
		" Konnichiwa, %v Yōkoso!",
	}

	randonIndex := rand.Intn(len(formats))
	messageFormat := formats[randonIndex]
	return messageFormat

}

func Hellos(names []string) (map[string]string, error) {
	// check if names is empty
	if len(names) == 0 {
		return nil, errors.New("empty names")
	}

	// A map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message

	}

	return messages, nil
}
