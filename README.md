# Dusk 🌑

Dusk is a minimal dependency Go library for calculating astronomical twilight, the lunar phase and the rise and set times of the moon and sun.

## Installation

Make sure you have Go installed ([download](https://golang.org/dl/)). Version `1.17` or higher is required for this package.

Initialize your project by creating a folder and then running `go mod init github.com/your/repo` ([learn more](https://blog.golang.org/using-go-modules)) inside the folder. Then install Dusk with the [`go get`](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) command:

```bash
go get -u github.com/observerly/dusk
```

## License

Dusk is free software licensed under the GNU General Public License v3.0 (GPL-3.0). See [LICENSE](./LICENSE).

The binary version of this program uses several open source libraries and components, which come with their own licensing terms. See below for an overview, and [LICENSE](./LICENSE) for details.

| Library attribution | License type |
|---------------------|--------------|
| [zsefvlol/timezonemapper](https://github.com/zsefvlol/timezonemapper) | MIT License |
