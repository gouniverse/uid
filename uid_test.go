package uid

import (
	"testing"
	"time"
)

func TestHumanUid(t *testing.T) {
	humanUid := HumanUid()
	humanUid2 := HumanUid()

	if humanUid == "" {
		t.Fatal("Human UID must not be null")
	}

	if len(humanUid) != 32 {
		t.Fatal("Human UID length must be 32 charcters")
	}

	if humanUid == humanUid2 {
		t.Fatal("Human UID and Human UID 2 must not be the same")
	}

	if humanUid > humanUid2 {
		t.Fatal("Human UID 1 must be smaller than Human UID 2")
	}
}

func TestMicroUid(t *testing.T) {
	microUid := MicroUid()
	microUid2 := MicroUid()

	if microUid == "" {
		t.Fatal("Micro UID must not be null")
	}

	if len(microUid) != 20 {
		t.Fatal("Micro UID length must be 20 charcters")
	}

	if microUid == microUid2 {
		t.Fatal("Micro UID and Micro UID 2 must not be the same")
	}

	if microUid > microUid2 {
		t.Fatal("Micro UID 1 must be smaller than Micro UID 2")
	}
}

func TestNanoUid(t *testing.T) {
	nanoUid := NanoUid()
	nanoUid2 := NanoUid()

	if nanoUid == "" {
		t.Fatal("Nano UID must not be null")
	}

	if len(nanoUid) != 23 {
		t.Fatal("Nano UID length must be 23 charcters")
	}

	if nanoUid == nanoUid2 {
		t.Fatal("Nano UID and Nano UID 2 must not be the same")
	}

	if nanoUid > nanoUid2 {
		t.Fatal("Nano UID 1 must be smaller than Nano UID 2")
	}
}

func TestSecUid(t *testing.T) {
	secUid := SecUid()
	time.Sleep(time.Second) // as its a seconds based ID we need at least a second between the generation
	secUid2 := SecUid()

	if secUid == "" {
		t.Fatal("Sec UID must not be null")
	}

	if len(secUid) != 14 {
		t.Fatal("Sec UID length must be 14 characters")
	}

	if secUid == secUid2 {
		t.Fatal("Sec UID and sec UID 2 must not be the same")
	}

	if secUid > secUid2 {
		t.Fatal("Sec UID 1 must be smaller than sec UID 2")
	}
}
