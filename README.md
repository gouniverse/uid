# Go Uid (Unique ID)

Go Uid generates unique identifying strings. Largest attention is paid on human friendly unique identifiers (dated digits).

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
    
