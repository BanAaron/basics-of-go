package aaron

func notExported() string {
	return "not exported"
}

func Exported() string {
	return "exported"
}
