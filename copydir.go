package misc

import (
	"io"
	"os"
)

func CopyFile(srcFile, destFile string) error {
	r, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	srcStat, err := os.Stat(srcFile)
	if err != nil {
		return err
	}
	return os.Chmod(destFile, srcStat.Mode())
}

func CopyDir(srcDir, destDir string) error {
	srcStat, err := os.Stat(srcDir)
	if err != nil {
		return err
	}
	if !srcStat.IsDir() {
		return Errorf("not a directory: %q", srcDir)
	}

	if Exists(destDir) {
		return Errorf("exists already: %q", destDir)
	}

	err = os.MkdirAll(destDir, srcStat.Mode())
	if err != nil {
		return err
	}

	srcF, err := os.Open(srcDir)
	if err != nil {
		return err
	}
	defer srcF.Close()

	srcFIs, err := srcF.Readdir(-1)
	if err != nil {
		return err
	}

	for _, srcFI := range srcFIs {
		srcDirExt := PathJoin(srcDir, srcFI.Name())
		destDirExt := PathJoin(destDir, srcFI.Name())

		if srcFI.IsDir() {
			err = CopyDir(srcDirExt, destDirExt)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcDirExt, destDirExt)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
