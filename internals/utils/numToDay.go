package utils

import "errors"

func NumToDay(dayNumb uint8) (string, error) {
	switch dayNumb {
	case 1:
		return "Senin", nil
	case 2:
		return "Selasa", nil
	case 3:
		return "Rabu", nil
	case 4:
		return "Kamis", nil
	case 5:
		return "Jumat", nil
	case 6:
		return "Sabtu", nil
	case 7:
		return "Minggu", nil
	default:
		return "", errors.New("hari tidak ditemukan")
	}
}
