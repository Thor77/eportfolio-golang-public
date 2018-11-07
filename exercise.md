Exercise: https://adventofcode.com/2016/day/5

Modifications:
* Your password has **4** characters
* The hexadecimal representation has to start with **3** zeros and it's **4th** character is a character of your password

Hints:
* Converting a string to an integer: [`strconv.Itoa`](https://golang.org/pkg/strconv/#Itoa)
* Calculate md5 hash of a `[]byte`: [`md5.Sum`](https://golang.org/pkg/crypto/md5/#Sum), hex-representation: [`hex.EncodeToString`](https://golang.org/pkg/encoding/hex/#EncodeToString)
* Checking if a string starts with a prefix: [`strings.HasPrefix`](https://golang.org/pkg/strings/#HasPrefix)
