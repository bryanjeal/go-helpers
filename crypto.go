// Copyright 2016 Bryan Jeal <bryan@jeal.ca>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

import (
	"crypto/rand"
	"crypto/sha512"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// Crypto type basically namespaces the helpers under the "crypto" section
var Crypto crypto

type crypto struct{}

// BCryptPasswordHasher takes a plaintext password, sha512 hashes it, and feeds it to bcrypt for further hashing.
// The password is sha512 hashed to "truncate" passwords longer than 72 characters (bcrypts max)
func (c crypto) BCryptPasswordHasher(password []byte) ([]byte, error) {
	hasher := sha512.New()
	hasher.Write(password)
	p, err := bcrypt.GenerateFromPassword(hasher.Sum(nil), 16)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// BCryptCompareHashPassword takes a plaintext password, sha512 hashes it, and feeds it to bcrypt for further hashing.
// The password is sha512 hashed to "truncate" passwords longer than 72 characters (bcrypts max)
func (c crypto) BCryptCompareHashPassword(hashed, password []byte) error {
	hasher := sha512.New()
	hasher.Write(password)
	err := bcrypt.CompareHashAndPassword(hashed, hasher.Sum(nil))
	if err != nil {
		return err
	}

	return nil
}

// GenerateRandomKey creates a random key with the given length in bytes.
func (c crypto) GenerateRandomKey(length int) ([]byte, error) {
	k := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil, err
	}
	return k, nil
}
