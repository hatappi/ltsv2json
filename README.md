# LTSV2JSON

convert from LTSV to JSON

## Installation

```sh
$ go get github.com/hatappi/ltsv2json
```

## Usage
use [sample file](./sample/fruits.ltsv)

**stdin**

```sh
$ cat sample/fruits.ltsv | ltsv2json
[{"name":"apple  num:3"},{"name":"banana num:5"},{"name":"orange num:2"}]
```

**Input File**

```
ltsv2json -i ./sample/fruits.ltsv
[{"name":"apple  num:3"},{"name":"banana num:5"},{"name":"orange num:2"}]
```


LTSV2JSON has features like `tail -f`

```sh
$ ltsv2json -i ./sample/fruits.ltsv -f
{"name":"apple  num:3"}
{"name":"banana num:5"}
{"name":"orange num:2"}

// another session
$ echo "name:grape\tnum:10" >> sample/fruits.ltsv

// It will be appended
$ ltsv2json -i ./sample/fruits.ltsv -f
{"name":"apple  num:3"}
{"name":"banana num:5"}
{"name":"orange num:2"}
{"name":"grape","num":10}
```

use [jq](https://stedolan.github.io/jq/)

```
$ cat ./sample/fruits.ltsv | ltsv2json |  jq -c '[.[] | .num] | add'
20
$ cat ./sample/fruits.ltsv | ltsv2json | jq -c '.[] | select(.name == "banana") | .num'
5
```

## Contributing

1. Fork it ( https://github.com/hatappi/ltsv2json/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
