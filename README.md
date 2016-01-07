# action-test

Testing tool for actions.

## Usage

Test actions

```sh
$ action-test
Usage of action-test:
  Application action-test runs tests from JSON or YAML files against actions from different sources.

Flags:
  -debug=false: Log debug info
  -sources="": Actions sources (comma separated directories & urls)
  -tests="": Files or directory containing YAML or JSON tests (can be glob pattern)
  -v=false: Verbose output: log all tests

Example YAML tests:

  -
    name: "filmweb.find"
    description: "Should find movie by title"
    arguments:
      title: "Pulp Fiction"
    expect:
      writers: "Quentin Tarantino"
      directors: "Quentin Tarantino"
      title: "Pulp Fiction"
      year: "1994"

$ action-test -tests=./example/tests.yaml -sources=file://example/actions
=== RUN filmweb.find

  Should find Pulp Fiction

    × title is empty (expected Pulp Fiction)
    × year is empty (expected 1994)
    ? description => Przemoc i odkupienie w opowieści o dwóch płatnych mordercach pracujących na zlecenie mafii, żon...
    ? countries => USA
    ? genres => Gangsterski
    ? poster => http://1.fwcdn.pl/po/10/39/1039/7517880.3.jpg
    × writers is empty (expected Quentin Tarantino)
    √ directors => Quentin Tarantino

--- FAIL filmweb.find (156.856401ms)
=== RUN filmweb.find

  Should find Apollo 13

    × year is empty (expected 1995)
    ? description => Podczas rutynowego lotu w kosmos na pokładzie statku Apollo 13 następuje wybuch.
    ? poster => http://1.fwcdn.pl/po/10/53/1053/7473756.3.jpg
    ? countries => USA
    ? genres => Dramat
    √ directors => Ron Howard
    × writers is empty (expected [William Broyles Jr. Al Reinert])
    × title is empty (expected Apollo 13)

--- FAIL filmweb.find (139.62857ms)
=== RUN filmweb.find

  Should find Avatar

    × year is empty (expected 2009)
    √ directors => James Cameron
    × writers is empty (expected James Cameron)
    ? poster => http://1.fwcdn.pl/po/91/13/299113/7332755.3.jpg
    ? description => Jake, sparaliżowany były komandos, otrzymuje misję i zostaje wysłany na planetę Pandora, gdzie ...
    ? genres => Sci-Fi
    ? countries => [USA Wielka Brytania]
    × title is empty (expected Avatar)

--- FAIL filmweb.find (128.831524ms)
--- FAIL   425.372557ms
```
