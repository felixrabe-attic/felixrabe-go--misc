package misc

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	// "runtime/debug"
)

func Contains(list []interface{}, elem interface{}) bool {
	for _, t := range list {
		if t == elem {
			return true
		}
	}
	return false
}

func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

func Exists(file string) bool {
	// see also http://stackoverflow.com/a/10510783/1034080
	_, err := os.Stat(file)
	return err == nil // !os.IsNotExist(err)
}

func Fatal(a ...interface{}) {
	// debug.PrintStack()
	fmt.Println(a...)
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	// debug.PrintStack()
	fmt.Printf(format+"\n", a...)
	os.Exit(1)
}

func PathJoin(a ...string) string {
	return filepath.Join(a...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func ThisDirN(n int) string {
	_, filename, _, ok := runtime.Caller(n)
	if !ok {
		Fatal("runtime.Caller() was not ok")
	}
	return filepath.Dir(filename)
}

func ThisDir() string {
	return ThisDirN(2)
}

func ThisDirJoin(a ...string) string {
	return filepath.Join(append([]string{ThisDirN(2)}, a...)...)
}
