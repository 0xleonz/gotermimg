# Gotermimg

`gotermin` is a small Go Library that displays images directly in your terminal,  
using the **Kitty/Sixel graphics protocol**.

## Build

From the project root:

```bash
go build -o bin/kitty-viewer ./cmd/kitty-viewer
```

Then run it:

```bash
./bin/kitty-viewer image.jpg
```

---


## Requirements

* Terminal emulator **Kitty** (or compatible with its graphics protocol).
* Go 1.18 or newer.

---

## How it works

Kitty’s graphics protocol uses ANSI escape sequences to transmit base64-encoded image data directly to the terminal.

## 🖼️ Example

If you’re using Kitty:

```bash
./bin/kitty-viewer ./testdata/cat.png
```

You should see the image rendered **inside your terminal**

---

## 📜 License

GPLv3 © 2025
