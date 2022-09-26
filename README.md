# Gonvif

A generated Onvif client in Go. Contains a simple command line utility to send Onvif
commands/queries to an Onvif compatible device. This also serves as an example of how to use the
generated client programatically.

## CLI Usage

Get the built `gonvif` binary for your system and architecture from our
[Release](https://github.com/eyetowers/gonvif/releases) page.

Then run a command of your choice from one of the supported Onvif ports, e.g., listing all device
profiles using the `media` port, providing the Onvif device URL and credentials:

```bash
gonvif -a http://IP[:PORT] -u USERNAME -p PASSWORD media get-profiles
```

### Shell completion

Get the shell completion script by running the following, using one of `bash`, `zsh`, `fish`,
`powershell`:

```bash
gonvif completion bash
```

## Client Usage

```golang
import (
    "log"

    "github.com/eyetowers/gonvif/pkg/client"
)

func main() {
    // Connect to the Onvif device.
    onvif, err := client.New("http://IP[:PORT]", "USERNAME", "PASSWORD")
    if err != nil {
        log.Fatal(err)
    }
    // Get the Media2 service client.
    media, err := onvif.Media2()
    if err != nil {
        log.Fatal(err)
    }
    // Make a request.
    resp, err := media.GetProfiles(&wsdl.GetProfiles{
        Type: []string{"All"},
    })
    if err != nil {
        log.Fatal(err)
    }
    // Process the response.
    log.Printf("Got profiles: %v", resp)
}
```

## License

Gonvif is open source, released under the [MIT license](./LICENSE).
