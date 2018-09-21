package cmd

//function for different language
func greeting(lang string) string {
	switch lang {
	case "pl":
		return "Czesc"
	case "es":
		return "Hola"
	default: //english
		return "Hello"
	}
}
