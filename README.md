# sqlt: More SQL types for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/soroushj/sqlt.svg)](https://pkg.go.dev/github.com/soroushj/sqlt)
[![CI](https://github.com/soroushj/sqlt/actions/workflows/ci.yml/badge.svg)](https://github.com/soroushj/sqlt/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/soroushj/sqlt/branch/main/graph/badge.svg?token=IABU4MSA7Y)](https://codecov.io/gh/soroushj/sqlt)
[![Go Report Card](https://goreportcard.com/badge/github.com/soroushj/sqlt)](https://goreportcard.com/report/github.com/soroushj/sqlt)

## Overview

**Note:** This module should not be used in new code. The generic type [`database/sql.Null`](https://pkg.go.dev/database/sql#Null) is available in the standard library since Go 1.22.

The sqlt package provides types that implement Go's [`database/sql.Scanner`](https://pkg.go.dev/database/sql#Scanner) and [`database/sql/driver.Valuer`](https://pkg.go.dev/database/sql/driver#Valuer) interfaces:

- [`github.com/soroushj/sqlt.NullRawMessage`](https://pkg.go.dev/github.com/soroushj/sqlt#NullRawMessage) - Nullable [`encoding/json.RawMessage`](https://pkg.go.dev/encoding/json#RawMessage)
- [`github.com/soroushj/sqlt.NullUUID`](https://pkg.go.dev/github.com/soroushj/sqlt#NullUUID) - Nullable [`github.com/google/uuid.UUID`](https://pkg.go.dev/github.com/google/uuid#UUID) (**Note:** This type should not be used in new code. [`github.com/google/uuid.NullUUID`](https://pkg.go.dev/github.com/google/uuid#NullUUID) is available since v1.3.0.)
