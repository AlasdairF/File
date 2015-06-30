package file

import (
 "os"
 "os/exec"
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
	err := os.Rename(oldpath, newpath)
	if err == nil {
		return nil
	}
	exec.Command(`mv`, oldpath, newpath).Run()
	if Exists(newpath) && !Exists(oldpath) {
		return nil
	}
	return err
}

func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	lst, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return lst, nil
}

func CountDir(dirname string) (int, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return 0, err
	}
	lst, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return 0, err
	}
	return len(lst), nil
}
