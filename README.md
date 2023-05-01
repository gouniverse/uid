# Go Uid (Unique ID) <a href="https://github.com/gouniverse/uid" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

[![Tests Status](https://github.com/gouniverse/uid/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/gouniverse/uid/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/uid)](https://goreportcard.com/report/github.com/gouniverse/uid)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/uid)](https://pkg.go.dev/github.com/gouniverse/uid)

THis package generates unique identifying strings. Largest attention is paid on human friendly unique identifiers (dated digits).

## Installation

```
go get -u https://github.com/gouniverse/uid
```

## Usage
 
 ```
 // HumainUid generates a UUID (32 digits) Format: YYYYMMDD-HHMM-SSMM-MMMMRRRRRRRRRRR
 humanUID := uid.HumanUid()
 
 // NanoUid generates a UID (23 digits) Format: YYYYMMDD-HHMMSS-MMMMMM-NNN
humanUID := uid.NanoUid()

// MicroUid generates a UID (20 digits) Format: YYYYMMDD-HHMMSS-MMMMMM
microidUID := uid.MicroUid()

// SecUid generates UID (14 digits) Format: YYYYMMDD-HHMMSS
secondsUID := uid.SecUid()
 ```

## Supported UID Types

It supports several types of unique identifiers. The type you may want to use depends on how random and long you want the identifier to be. For most of the user cases a Micro UUID (20 chars) should be fine. A human UUID (32 chars) should almost never be used where a human is involved as too "mind bogling" to work with.

1. Human UUID (32 digits)

    Format: YYYYMMDD-HHMM-SSMM-MMMMRRRRRRRRRRR

    2017111908492665991498485465 (with dashes: 20171119-0849-2665-991498485465)

2. Nano UID (23 digits)

    Format: YYYYMMDD-HHMMSS-MMMMMM-NNN

    Examples:

    20171119084926659914984 (with dashes: 20171119-084926-659914-984)

3. Micro UID (20 digits)

    Format: YYYYMMDD-HHMMSS-MMMMMM

    Examples:

    20171119084926659914 (with dashes: 20171119-084926-659914)

4. Seconds UID (14 digits)

    Format: YYYYMMDD-HHMMSS

    Examples:

    20171119084926 (with dashes: 20171119-084926)
    

## Change Log

2021.12.19 - Master branch changed to main
2021.12.19 - Added tests

## Similar Packages

- https://github.com/jaevor/go-nanoid (random 21 characters)
- https://github.com/zheng-ji/goSnowFlake (timestamp-workerid-sequence)
- https://github.com/damdo/randid (random IDs)
- https://github.com/matthewmueller/uid.go (shorcodes)
- https://github.com/aohorodnyk/uid (random IDs)
- https://github.com/google/uuid (UUIDs)
- https://github.com/oklog/ulid (???)
- https://github.com/chilts/sid (setial IDs)
- https://datatracker.ietf.org/doc/html/draft-ietf-uuidrev-rfc4122bis (GUID 6?)
