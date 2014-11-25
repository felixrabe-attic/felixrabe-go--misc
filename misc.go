package misc

import (
	"fmt"
	"io"
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

func FatalIf(err error) {
	if err != nil {
		Fatal(err)
	}
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
}

func OpenOrCreate(name string) (file *os.File, err error) {
	f, err := os.Open(name)
	if err == nil {
		// no error, all good
		return f, nil
	}
	if !os.IsNotExist(err) {
		// err does not indicate that the file does not exist
		return nil, err
	}
	// file does not exist, try to create it
	f, err = os.Create(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func PathJoin(a ...string) string {
	return filepath.Join(a...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
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
