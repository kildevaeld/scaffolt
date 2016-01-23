package engine

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/kildevaeld/scaffolt"
)

var FileAlreadyExistsErr = errors.New("File already exists")

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func IsDir(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return stat.IsDir()
	}
	return false
}

func IsFile(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return !stat.IsDir()
	}
	return false
}

func Interpolate(name, str string, ctx scaffolt.Context) (string, error) {
	t, err := template.New(name).Parse(str)

	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, ctx.Locals())
	if err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}

func InterpolateOrDefault(name, str string, ctx scaffolt.Context) string {
	out, err := Interpolate(name, str, ctx)
	if err != nil {
		return str
	}
	return out
}

func InterpolateFile(path string, ctx scaffolt.Context) (string, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	base := filepath.Base(path)

	return Interpolate(base, string(bs), ctx)
}

func createFile(dst, src string, overwrite bool) (err error) {

	if Exists(dst) && !overwrite {
		return FileAlreadyExistsErr
	}

	dir := filepath.Dir(dst)

	if !Exists(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	var t, s *os.File
	if t, err = os.Open(dst); err != nil {
		return err
	}
	defer t.Close()
	if s, err = os.Open(src); err != nil {
		return err
	}
	defer s.Close()
	if _, err := io.Copy(t, s); err != nil {
		return err
	}

	return nil
	/*e := ioutil.WriteFile(fullpath, content, 0755)

	if e != nil {
		return e
	}*/

}

func createFileBytes(dst string, content []byte, overwrite bool) (err error) {
	if Exists(dst) && !overwrite {
		return FileAlreadyExistsErr
	}

	dir := filepath.Dir(dst)

	if !Exists(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return ioutil.WriteFile(dst, content, 0755)
}

func stringToBoolean(str string) bool {
	l := strings.ToLower(str)
	if l == "y" || l == "j" || l == "true" || l == "yes" {
		return true
	}
	return false
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
