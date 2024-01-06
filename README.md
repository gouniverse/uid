# UID (Unique ID) <a href="https://gitpod.io/#https://github.com/gouniverse/uid" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

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
 // HumanUid generates a UUID (32 digits) Format: YYYYMMDD-HHMM-SSMM-MMMMNNNRRRRRRRRR
 humanUID := uid.HumanUid()
 
 // NanoUid generates a UID (23 digits) Format: YYYYMMDD-HHMMSS-MMMMMM-NNN
humanUID := uid.NanoUid()

// MicroUid generates a UID (20 digits) Format: YYYYMMDD-HHMMSS-MMMMMM
microUID := uid.MicroUid()

// SecUid generates UID (14 digits) Format: YYYYMMDD-HHMMSS
secondsUID := uid.SecUid()
 ```

## Supported UID Types

It supports several types of unique identifiers. 

The type you want to use will usually depends on two considerations:

1. How random you want it to be? The longer the identifier, the more the chances of collision reduce
2. How long you want the identifier to be? The longer the identifier, reduces the readability, as well as the storage space to store it.

For most of the user cases a Micro UUID (20 chars) should be fine. A human UUID (32 chars) should be avoided where a human is involved as too "mind bogging" to work with.

1. Human UUID (32 digits)

    Format: YYYYMMDD-HHMM-SSMM-MMMMNNNRRRRRRRRR

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

5. Timestamp (10 digits)
    Unit timestamp, seconds precision

    Format: 1234567890

    Examples:

    1704524414


6. TimestampMicro (16 digits)
    Unit timestamp, microseconds precision

    Format: 1234567890123456

    Examples:

    1704524414548721

6. TimestampNano (19 digits)
    Unit timestamp, nanoseconds precision

    Format: 1234567890123456789

    Examples:

    1704524414548721308

7. Uuid (32 characters)
    Random V4 UUID. UUID (Universally Unique IDentifier), also known as GUID (Globally Unique IDentifier)

    Format: abcdef1234567890abcdef1234567890

    Examples:

    459e2999bd071151a23d643da42c2cc2
    

## Change Log
2024.01.06 - Added Timestamp and Uuid functions
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
- https://github.com/chilts/sid (serial IDs)
- https://datatracker.ietf.org/doc/html/draft-ietf-uuidrev-rfc4122bis (GUID 6?)
