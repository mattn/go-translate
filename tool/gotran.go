package main

import (
	"google/language/translate"
	"os"
	"flag"
)

func main() {
	from := flag.String("f", "en", "translate from");
	to := flag.String("t", "ja", "translate to");
	flag.Parse();
	s, err := translate.Translate(*from, *to, flag.Arg(0))
	if err != nil {
		os.Stderr.WriteString(err.String() + "\n")
	} else {
		println(s);
	}
}
