# metapkg

NOTE: This project is currently only an experiment, not a usable tool.

## Use case

Your do not install your software using a system package, but your software depends on (and/or conflicts with) system packages. You can install a "meta" package that defines dependencies and conflicts, but includes no artifacts or scripts.

## Example

```console
> go build
> ./metapkg > meta.rpm
> rpm -i meta.rpm
```

```console
error: Failed dependencies:
	foo >= 1.0 is needed by meta-0:1.0.0-1.x86_64
	foo < 2.0 is needed by meta-0:1.0.0-1.x86_64
```
