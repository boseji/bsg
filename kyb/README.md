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

## `kyb` – Cross-Platform Keyboard Input Simulation for Go

`kyb` is a small, OS-aware Go package for **simulating keyboard input** on
**Windows, Linux, and macOS** using a **single, unified API**.

It automatically selects the best available backend per operating system while
keeping your application code clean and portable.

### Features for the `kyb` Package

- Type strings programmatically
- Send individual key presses
- Configurable typing delay (human-like typing)
- Same API on Windows, Linux, and macOS
- OS-specific backends selected via Go build tags
- **No CGO required** (default configuration)
- Dependency checks where applicable

### Supported Keys in `kyb`

#### Common Keys

* `enter`
* `tab`
* `esc`
* `space`
* `shift`
* `ctrl`
* `alt`

> Support varies slightly by OS backend.

### Limitations of `kyb`

* Wayland (Linux) does not allow synthetic input
* macOS blocks input without Accessibility permission
* Games and secure applications may ignore synthetic input
* Unicode and complex layouts are limited on non-native backends

### Extensibility of `kyb`

The package is designed to be extended with:

* Modifier combos (`Ctrl+C`, `Cmd+V`)
* Randomized delay (human typing)

## Basic Usage of `kyb` Package

```go
import "kyb"

// Type text using default delay
kyb.Type("Hello world")

// Enable human-like typing
kyb.SetDelay(75 * time.Millisecond)
kyb.Type("Typing slowly...")

// Disable delay for one call
kyb.TypeFast("Instant typing")

// Press a single key
kyb.KeyPress("enter")
```

## Typing Delay in `kyb`

Typing delay is **globally configurable** and applied consistently across OSes.

```go
kyb.SetDelay(50 * time.Millisecond)
kyb.Type("Human-like typing")
```

| OS      | Delay implementation            |
| ------- | ------------------------------- |
| Linux   | Native `xdotool --delay`        |
| Windows | `time.Sleep` between keystrokes |
| macOS   | `time.Sleep` between keystrokes |

## Availability Check built into `kyb`

```go
if !keyboard.Available() {
    log.Fatal("keyboard input simulation not available on this system")
}
```

This helps to make sure that that required packages are installed for
the OS platform.

----

## Linux Support for `kyb`

### Backend

**`xdotool` (X11 only - No Wayland)**

### Requirements

```bash
# for Debian / Ubuntu
sudo apt install xdotool
# or for Arch based
sudo pacman -S xdotool
# or for Fedora
sudo dnf install xdotool
```

### Display Server IMPORTANT

* ✅ **X11 supported**
* ❌ **Wayland not supported**

Check your session type:

```bash
echo $XDG_SESSION_TYPE
```

### Notes

* The package automatically checks for `xdotool`
* Returns `ErrNotSupported` if missing or unsupported
* Window targeting and modifiers are supported via xdotool

----

## macOS Support for `kyb`

### Backend

* **AppleScript** via `osascript`
* Uses **System Events**

### Requirements

* macOS 11+
* `osascript` (included by default)

### Accessibility Permission (Required) IMPORTANT

Your application **must be allowed** under:

```
System Settings → Privacy & Security → Accessibility
```

Without this permission:

* Keystrokes will be silently ignored
* No explicit error is returned by macOS

### Notes

* Typing is done character-by-character
* Delay is handled in Go (not AppleScript)
* Slightly slower than native Quartz events

----

## Windows Support for `kyb`

### Backend

* Native **Win32 `SendInput` API**
* Uses `user32.dll`
* No external dependencies

### Requirements

* Windows 10 or later
* No additional installs required

### Notes

* Input is sent to the **currently focused window**
* Elevated apps can only be controlled by elevated processes
* Secure desktops (UAC prompts, login screen) are not supported

----

## Summary of `kyb` Implementation Support

| OS      | Backend         | External Dependency | Notes                             |
| ------- | --------------- | ------------------- | --------------------------------- |
| Windows | Win32 SendInput | ❌                  | Native, fast                      |
| Linux   | xdotool         | ✅                  | X11 only                          |
| macOS   | AppleScript     | ❌                  | Accessibility permission required |

----

## License

This project is released under the GNU General Public License v2. See the [LICENSE](../LICENSE.txt) file for details.

Sources: <https://github.com/boseji/bsg>

`bsg` - Boseji's Security and Privacy Utilities.

Copyright (C) 2025-2026 by Abhijit Bose (aka. Boseji)

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



