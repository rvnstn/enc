package cmd

import (
	"bytes"
	"fmt"
	"io"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

func ReadKey(r io.Reader) (*crypto.Key, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read from stdin: %v", err)
	}
	isArmored := bytes.HasPrefix(data, []byte("-----BEGIN PGP PRIVATE KEY BLOCK-----"))
	if !isArmored {
		isArmored = bytes.HasPrefix(data, []byte("-----BEGIN PGP PUBLIC KEY BLOCK-----"))
	}
	if isArmored {
		key, err := crypto.NewKeyFromArmored(string(data))
		if err != nil {
			return nil, fmt.Errorf("unarmor key: %v", err)
		}
		return key, nil
	} else {
		key, err := crypto.NewKey(data)
		if err != nil {
			return nil, fmt.Errorf("parse key: %v", err)
		}
		return key, nil
	}
}