# remo-numpad

## Description

remo-numpad enable you to controll Home Appliances and Devices via Nature Remo with the Numeric KeyPad.

<img src="https://github.com/jedipunkz/remo-numeric-keypad/blob/master/pix/numeric_key_with_raspberyy_pi.jpg" width=512>

## Pre-Requirement(s)

- gobot.io with My PR which add symbol key events: https://github.com/hybridgroup/gobot/pull/767

## Installation

```bash
$ go get github.com/jedipunkz/remo-numpad.git
```

## Setup `$HOME/.remo-key.yaml` file.

This repository includes example yaml file. see `example_yamls` dir.

## Boot Process

```bash
$ remo-numpad server
```

## Scan Pressed Key and Print Key Char, Key Code

```bash
$ remo-numpad scan # then if you press any key, key character and code will be printed.
```
