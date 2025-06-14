>
> ॐᳬ᳞ भूर्भुवः स्वः
>
> तत्स॑वि॒तुर्वरे॑ण्यं॒
>
> भर्गो॑ दे॒वस्य॑ धीमहि।
>
> धियो॒ यो नः॑ प्रचो॒दया॑त्॥
>

#  बी.स.जी - बोसजी के द्वारा रचित सुरक्षा एवं गोपनीयता हेतु तन्त्राक्ष्।

> एक सुरक्षा एवं गोपनीयता केंद्रित तंत्राक्षों का संकलन।

***एक रचनात्मक भारतीय उत्पाद ।***

## bsg - Boseji's Security and Privacy Utilities

A collection of Security and Privacy utilities and some notes for help.

This is **Golang** package collection as well as few utility command line programs.

## `gen` Package - part of the `bsg` project

A drop-in cryptographically secure replacement for math/rand/v2 in Go, using the crypto/rand package.

This package provides secure alternatives for generating random integers, floating-point numbers, permutations, and shuffling — suitable for use in cryptographic or security-sensitive applications.

### Features

* Cryptographically secure versions of:

  * `Int`, `IntN`, `Int32`, `Int32N`, `Int64`, `Int64N`
  * `Uint`, `UintN`, `Uint32`, `Uint32N`, `Uint64`, `Uint64N`
  * `Float32`, `Float64`
  * `Shuffle`, `Perm`

* Random Length String Generation

  * Supplied with Default charset `gen.CharSet`

    `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`

### Usage Examples

```go
import "bsg/gen"

n := gen.IntN(100)      // int in [0, 100)
f := gen.Float64()      // float64 in [0.0, 1.0)
perm := gen.Perm(10)    // secure permutation of [0..9]
```

### API Reference

| Function         | Description                                                                 |
| ---------------- | --------------------------------------------------------------------------- |
| `Int()`          | Secure random `int`                                                         |
| `IntN(n int)`    | Secure random `int` in `[0, n)`                                             |
| `Int32()`        | Secure random `int32`                                                       |
| `Int32N(n)`      | Secure random `int32` in `[0, n)`                                           |
| `Int64()`        | Secure random `int64`                                                       |
| `Int64N(n)`      | Secure random `int64` in `[0, n)`                                           |
| `Uint()`         | Secure random `uint`                                                        |
| `UintN(n)`       | Secure random `uint` in `[0, n)`                                            |
| `Uint32()`       | Secure random `uint32`                                                      |
| `Uint32N(n)`     | Secure random `uint32` in `[0, n)`                                          |
| `Uint64()`       | Secure random `uint64`                                                      |
| `Uint64N(n)`     | Secure random `uint64` in `[0, n)`                                          |
| `Float32()`      | Secure random `float32` in `[0.0, 1.0)`                                     |
| `Float64()`      | Secure random `float64` in `[0.0, 1.0)`                                     |
| `Shuffle(n, fn)` | Securely shuffles a range using `swap(i, j)`                                |
| `Perm(n)`        | Returns secure permutation of `[0, n)`                                      |
| `String(ch, n)`  | Securely generate a string of Random items from the supplied character-set. |


## License

This project is released under the GNU General Public License v2. See the [LICENSE](../LICENSE.txt) file for details.

Sources: <https://github.com/boseji/bsg>

`bsg` - Boseji's Security and Privacy Utilities.

Copyright (C) 2025 by Abhijit Bose (aka. Boseji)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License version 2 only
as published by the Free Software Foundation.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

SPDX-License-Identifier: `GPL-2.0-only`

Full Name: `GNU General Public License v2.0 only`

Please visit <https://spdx.org/licenses/GPL-2.0-only.html> for details.

