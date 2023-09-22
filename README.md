# remo-numpad

## Description

remo-numpad enable you to controll Home Appliances and Devices via Nature Remo with the Numeric KeyPad.

<img src="https://github.com/jedipunkz/remo-numeric-keypad/blob/master/pix/numeric_key_with_raspberyy_pi.jpg" width=512>

## Installation

```bash
$ go install github.com/jedipunkz/remo-numpad@latest
```

## Setup `$HOME/.remo-numpad.yaml` file.

This repository includes example yaml file. see `example_yamls` dir.

## Boot Process

```bash
$ remo-numpad server
```

## Scan Pressed Key and Print Key Char, Key Code

```bash
$ remo-numpad scan # then if you press any key, key character and code will be printed.
```
