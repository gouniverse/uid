package uid

import (
	"crypto/rand"
	"strings"
	"time"
)

// HumainUid generates a UUID (32 digits) Format: YYYYMMDD-HHMM-SSMM-MMMMRRRRRRRRRRR
func HumanUid() string {
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:32]
}

// NanoUid generates a UID (23 digits) Format: YYYYMMDD-HHMMSS-MMMMMM-NNN
func NanoUid() string {
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:23]
}

// MicroUid generates a UID (20 digits) Format: YYYYMMDD-HHMMSS-MMMMMM
func MicroUid() string {
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:20]
}

// SecUid generates UID (14 digits) Format: YYYYMMDD-HHMMSS
func SecUid() string {
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:14]
}
