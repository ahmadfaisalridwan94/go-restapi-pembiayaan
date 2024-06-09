package dates

import "strings"

func ConvertToID(p string) string {
	r := strings.NewReplacer(
		"January", "Januari",
		"February", "Februari",
		"March", "Maret",
		"April", "April",
		"May", "Mei",
		"June", "Juni",
		"July", "Juli",
		"August", "Agustus",
		"September", "September",
		"October", "Oktober",
		"November", "November",
		"December", "Desember",
		"Sunday", "Minggu",
		"Monday", "Senin",
		"Tuesday", "Selasa",
		"Wednesday", "Rabu",
		"Thursday", "Kamis",
		"Friday", "Jumat",
		"Saturday", "Sabtu",
	)

	return r.Replace(p)
}
