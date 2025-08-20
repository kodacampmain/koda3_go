package intermediate

func ChangeName(oldName *string) {
	if *oldName == "Andi" {
		*oldName = "Budi"
		return
	}
	*oldName = "Candra"
}
