package file

import "os"

func Exists(f string) bool {
	if _, err := os.Stat(f); err == nil {
		return true
	} else {
		return false
	}
}

func Size(f string) (int, error) {
	v, err := os.Stat(f)
	if err != nil {
		return 0, err
	}
	return int(v.Size()), nil
}
