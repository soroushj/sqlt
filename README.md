# sqlt

More SQL types for Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/soroushj/sqlt)](https://goreportcard.com/report/github.com/soroushj/sqlt)
[![Build Status](https://travis-ci.org/soroushj/sqlt.svg?branch=master)](https://travis-ci.org/soroushj/sqlt)
[![codecov](https://codecov.io/gh/soroushj/sqlt/branch/master/graph/badge.svg)](https://codecov.io/gh/soroushj/sqlt)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/soroushj/sqlt)](https://pkg.go.dev/github.com/soroushj/sqlt)

## Overview

The sqlt package provides types that implement Go's [`sql.Scanner`](https://pkg.go.dev/database/sql#Scanner) and [`driver.Valuer`](https://pkg.go.dev/database/sql/driver#Valuer) interfaces:

- [`sqlt.NullRawMessage`](https://pkg.go.dev/github.com/soroushj/sqlt#NullRawMessage) - Nullable [`json.RawMessage`](https://pkg.go.dev/encoding/json/#RawMessage)
- [`sqlt.NullUUID`](https://pkg.go.dev/github.com/soroushj/sqlt#NullUUID) - Nullable [`uuid.UUID`](https://pkg.go.dev/github.com/google/uuid#UUID)
