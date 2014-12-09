package file

import (
 "os"
 "os/exec"
 "errors"
)

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

func Remove(f string) error {
	return os.Remove(f)
}

func Move(oldpath, newpath string) error {
	if err := os.Rename(oldpath, newpath); err == nil {
		return nil
	}
	exec.Command(`mv`, oldpath, newpath).Run()
	if Exists(newpath) && !Exists(oldpath) {
		return nil
	}
	return errors.New(`Unable to move ` + oldpath)
}
