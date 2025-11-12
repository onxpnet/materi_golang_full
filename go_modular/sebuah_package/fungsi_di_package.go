package sebuah_package

// fungsi dengan awalan CAPITAL akan tereksport
func SebuahFungsiDiPackage(say string) string {
	return "fungsi di package bilang: " + say
}

// fungsi internal yang tidak akan tereksport awalan Huruf KECIL
func sebuahFungsi(say string) string {
	return "fungsi internal: " + say
}
