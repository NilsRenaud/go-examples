package gotests


func MyExportedFunction(name string) string {
	if name != "" {
		return "OK"
	} else {
		return "KO"
	}
}
