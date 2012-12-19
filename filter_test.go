package filter

import (
	"bytes"
	"io/ioutil"
	"testing"

	"crypto/cipher"
	"code.google.com/p/go.crypto/blowfish"
)

func TestEncrypt(t *testing.T) {

	key := []byte("This is Key!")
	c, err := blowfish.NewCipher(key)
	if err != nil {
		t.Error(err.Error())
	}

	str := "Hello Filter!"
	buf := bytes.NewBuffer(nil)
	w := &cipher.StreamWriter {
		cipher.NewCFBEncrypter(c, make([]byte, c.BlockSize())),
		buf,
		nil,
	}
	encryptFilter := &Filter{buf, w}
	encryptFilter.Write([]byte(str))

	encrypted, _ := ioutil.ReadAll(encryptFilter)
	buf2 := bytes.NewBuffer(nil)
	r := &cipher.StreamReader {
		cipher.NewCFBDecrypter(c, make([]byte, c.BlockSize())),
		buf2,
	}
	decryptFilter := &Filter{r, buf2}
	decryptFilter.Write(encrypted)

	decrypted, _ := ioutil.ReadAll(decryptFilter)

	if str != string(decrypted) {
		t.Error("Expect", str, "but actual", string(decrypted))
	}
}
