# lojban

The program converts arabic numerals to lojban numerals and backward.

>   Lojban (from logji and bangu, or logic language) is a syntactically unambiguous
>   constructed language based on predicate logic, created in 1987 by The Logical
>   Language Group. Lojban is derivated from Loglan, a language invented in 1955 by
>   James Cooke Brown who claimed copyright on it, so Lojban started afresh from
>   its lexical basis to create a whole new vocabulary.

http://www.languagesandnumbers.com/how-to-count-in-lojban/en/jbo/
https://rationalwiki.org/wiki/Lojban

# Usage

From arabic to lojban:

```bash
$ lojban 123
pareci

$ lojban 954
somuvo
```

From arabic to lojban with prefix `man`:

```bash
$ lojban 123 -p man
manpareci

$ lojban 954 -p man
mansomuvo
```

From lojban to arabic:

```bash
$ lojban pareci
123

$ lojban sopaxa
916
```

From lojban to arabic with prefix `man`:

```bash
$ lojban manpareci -p man
123

$ lojban mansopaxa -p man
916
```

# Installation

go-get:

```
go get github.com/kovetskiy/lojban
```

archlinux:
```
yaourt -Sy lojban-git
```
