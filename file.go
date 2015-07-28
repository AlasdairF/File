package file

import (
 "os"
 "os/exec"
 "io"
 "fmt"
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

func Copy(src, dst string) (err error) {
    sfi, err := os.Stat(src)
    if err != nil {
        return
    }
    if !sfi.Mode().IsRegular() {
        // cannot copy non-regular files (e.g., directories,
        // symlinks, devices, etc.)
        return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
    }
    dfi, err := os.Stat(dst)
    if err != nil {
        if !os.IsNotExist(err) {
            return
        }
    } else {
        if !(dfi.Mode().IsRegular()) {
            return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
        }
        if os.SameFile(sfi, dfi) {
            return
        }
    }
    if err = os.Link(src, dst); err == nil {
        return
    }
	// copyFileContents
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
}
