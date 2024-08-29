## Terminal password generator

### Install

```sh
go install github.com/enslerman/passgen@latest
```

### Usage

```sh
# passgen -l [password len] -c [copy to clipboard] -s [use spec symbols]
passgen
# Success generate password with len: 64; spec symbols: false; copy to clipboard: true

passgen -l 16 -s true -p true
# Success generate password with len: 16; spec symbols: true; copy to clipboard: true
# x1P$M7gAHED_044P

passgen -l 16 -s true -c false -p true
# Success generate password with len: 16; spec symbols: true; copy to clipboard: false
# i#HBwIlo3OB37dXY
```

Output

```
Success generate password with len: 64; spec symbols: false; copy to clipboard: true
zwB1utD5OEIEWvvoTA4N1UAflPiY8GWhw6w3PJGCDDO5ophfR2AZqfFRL4EqiQaC
```
