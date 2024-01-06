package uid

import (
	"crypto/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// HumanUid generates a UUID (32 digits) Format: YYYYMMDD-HHMM-SSMM-MMMMNNNRRRRRRRRR
func HumanUid() string {
	time.Sleep(1 * time.Nanosecond)
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:32]
}

// NanoUid generates a UID (23 digits) Format: YYYYMMDD-HHMMSS-MMMMMM-NNN
func NanoUid() string {
	time.Sleep(time.Nanosecond) // as its a nanoseconds based ID we need at least a nanosecond between the generations to avoid collisions
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:23]
}

// MicroUid generates a UID (20 digits) Format: YYYYMMDD-HHMMSS-MMMMMM
func MicroUid() string {
	time.Sleep(time.Microsecond) // as its a microseconds based ID we need at least a microsecond between the generations to avoid collisions
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:20]
}

// SecUid generates UID (14 digits) Format: YYYYMMDD-HHMMSS
func SecUid() string {
	time.Sleep(time.Second) // as its a seconds based ID we need at least a second between the generations to avoid collisions
	r, _ := rand.Prime(rand.Reader, 64)
	id := time.Now().UTC().Format("20060102150405.0000000")
	id = strings.ReplaceAll(id, ".", "")
	id += r.String()
	return id[0:14]
}

func Timestamp() string {
	time.Sleep(time.Second) // as its a seconds based ID we need at least a second between the generations to avoid collisions
	now := time.Now().UTC().Unix()
	return strconv.FormatInt(now, 10)
}

func TimestampMicro() string {
	time.Sleep(time.Microsecond) // as its a microseconds based ID we need at least a microsecond between the generations to avoid collisions
	now := time.Now().UTC().UnixMicro()
	return strconv.FormatInt(now, 10)
}

func TimestampNano() string {
	time.Sleep(time.Nanosecond) // as its a nanoseconds based ID we need at least a nanosecond between the generations to avoid collisions
	now := time.Now().UTC().UnixNano()
	return strconv.FormatInt(now, 10)
}

func Uuid() string {
	return strings.ReplaceAll(UuidFormatted(), "-", "")
}

func UuidFormatted() string {
	return uuid.New().String()
}

func UuidV1() string {
	return strings.ReplaceAll(UuidV1Formatted(), "-", "")
}

func UuidV1Formatted() string {
	return uuid.Must(uuid.NewUUID()).String()
}

func UuidV3(namespace string, data []byte) (string, error) {
	uid, err := UuidV3Formatted(namespace, data)

	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(uid, "-", ""), nil
}

func UuidV3Formatted(namespace string, data []byte) (string, error) {
	id, err := uuid.FromBytes([]byte(namespace))

	if err != nil {
		return "", err
	}

	return uuid.Must(uuid.NewMD5(id, data), nil).String(), nil
}

func UuidV4() string {
	return strings.ReplaceAll(UuidV4Formatted(), "-", "")
}

func UuidV4Formatted() string {
	return uuid.New().String()
}

func UuidV5(namespace string, data []byte) (string, error) {
	uid, err := UuidV5Formatted(namespace, data)

	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(uid, "-", ""), nil
}

func UuidV5Formatted(namespace string, data []byte) (string, error) {
	id, err := uuid.FromBytes([]byte(namespace))

	if err != nil {
		return "", err
	}

	return uuid.NewSHA1(id, data).String(), nil
}

func UuidV6() string {
	return strings.ReplaceAll(UuidV6Formatted(), "-", "")
}

func UuidV6Formatted() string {
	return uuid.Must(uuid.NewV6()).String()
}

func UuidV7() string {
	return strings.ReplaceAll(UuidV7Formatted(), "-", "")
}

func UuidV7Formatted() string {
	return uuid.Must(uuid.NewV7()).String()
}
