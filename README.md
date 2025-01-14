[![test](https://github.com/kkentzo/ut/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/kkentzo/ut/actions/workflows/test.yml)

# ut

A utility that converts its arguments from unix timestamps to RFC3339 timestamps.

The program auto-detects the timestamp resolution; supported resolutions include:

- seconds
- milliseconds
- microseconds
- nanoseconds

Usage:

```bash
$ ut TIMESTAMP_1 TIMESTAMP_2 ... TIMESTAMP_N
```

Example:

```bash
$ ut 1736850409254
2025-01-14T12:26:49+02:00
```
