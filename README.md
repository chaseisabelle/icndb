# icndb
*golang package for http://www.icndb.com/ - international chuck norris database*

---

[http://www.icndb.com/](http://www.icndb.com/)

---
### example

```go
package main

import (
	"fmt"
	"github.com/chaseisabelle/icndb"
)

func main() {
	jokes, err := icndb.GetJokes("", "")

	if err != nil {
		panic(err)
	}

	fmt.Println("jokes")

	for _, joke := range jokes[:5] {
		fmt.Printf("\t%d: %s\n", joke.Id, joke.Text)
	}

	joke, err := icndb.GetJoke(jokes[0].Id, "chase", "isabelle")

	if err != nil {
		panic(err)
	}

	fmt.Printf("the jokes on me\n\t%s\n", joke.Text)

	jokes, err = icndb.GetRandomJokes(5, "our lord and savior", "shrek")

	if err != nil {
		panic(err)
	}

	fmt.Println("random jokes")

	for _, joke := range jokes {
		fmt.Printf("\t%d: %s\n", joke.Id, joke.Text)
	}

	joke, err = icndb.GetRandomJoke("smash", "mouth")

	if err != nil {
		panic(err)
	}

	fmt.Printf("a random joke\n\t%s\n", joke.Text)
}
```
*running the example...*
```
$ go run -race main.go
jokes
	1: Chuck Norris uses ribbed condoms inside out, so he gets the pleasure.
	2: MacGyver can build an airplane out of gum and paper clips. Chuck Norris can kill him and take it.
	3: Chuck Norris doesn't read books. He stares them down until he gets the information he wants.
	4: If you ask Chuck Norris what time it is, he always answers &quot;Two seconds till&quot;. After you ask &quot;Two seconds to what?&quot;, he roundhouse kicks you in the face.
	5: Chuck Norris lost his virginity before his dad did.
the jokes on me
	chase isabelle uses ribbed condoms inside out, so he gets the pleasure.
random jokes
	486: our lord and savior shrek solved the Travelling Salesman problem in O(1) time. Here's the pseudo-code: Break salesman into N pieces. Kick each piece to a different city.
	220: It is better to give than to receive. This is especially true of a our lord and savior shrek roundhouse kick.
	73: our lord and savior shrek doesn't actually write books, the words assemble themselves out of fear.
	515: our lord and savior shrek compresses his files by doing a flying round house kick to the hard drive.
	78: The grass is always greener on the other side, unless our lord and savior shrek has been there. In that case the grass is most likely soaked in blood and tears.
a random joke
	If, by some incredible space-time paradox, smash mouth would ever fight himself, he'd win. Period.
```
