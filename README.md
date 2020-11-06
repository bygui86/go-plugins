
# go-plugins

Sample project to explore plugins offered by Golang

#### Pros

- Go plugins give you a tool to help you adopt the single responsibility principle and separation of concerns.
- It helps to break down your code into small manageable and reusable components.
- It gives you a way to load plugins dynamically during application runtime without recompiling the program.

#### Cons

- The environment for building a Go plugin such as OS, Go language version, and dependency versions must match exactly.
- As of now, unloading plugins are not allowed without restarting the program.
- You cannot replace a plugin with a newer version during runtime; that’s because Go doesn’t currently support unloading plugins.
- As of Go v1.11, you can only build plugins on Linux, FreeBSD, and Mac.

#### Example program

The basic shipping calculator will give you rates based on which shipping method and parcel weight you provide. You can support different shipping methods by adding new plugins, and the calculator will produce the rates and currency based on your preferred shipping method.

## Run

Calculate using fedex
```shell script
make run-fedex
```

Calculate using royalmail
```shell script
make run-royalmail
```

## Links

- https://medium.com/swlh/how-to-build-extensible-go-applications-with-plugins-97ad657f9a62
- https://github.com/ManiMuridi/go-plugins-shipping-calculator
- https://github.com/olekukonko/tablewriter
