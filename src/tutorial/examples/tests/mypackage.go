package gotests


func MyExportedFunction(name string) string {
	if name != "" {
		return "OK"
	} else {
		return "KO"
	}
}

func myPrivateFunction(name string) string {
	if name != "" {
		return "KO" //bug
	} else {
		return "KO"
	}
}
