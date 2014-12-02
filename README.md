`tros` provides methods to sort slices of structs using reflection

[![GoDoc](https://godoc.org/github.com/ImJasonH/tros?status.svg)](https://godoc.org/github.com/ImJasonH/tros)
[![Build Status](https://travis-ci.org/ImJasonH/tros.svg?branch=master)](https://travis-ci.org/ImJasonH/tros)

Sorting Idiomatically in Go
-----

```
type Thing struct {
	Name           string
	Weight, Height int
	Awesome        bool
}

func sortThingsByName(things []Thing) {
	sort.Sort(thingWrapper)
}

type thingWrapper []Thing

func (w thingWrapper) Len() int           { return len(w) }
func (w thingWrapper) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w thingWrapper) Less(i, j int) bool { return w[i].Name < w[j].Name }

// TODO: Support sorting by weight, height, awesomeness, etc.
```

Sorting slices of Go structs idiomatically is easy, if you only intend to do it for a few different fields, and you know what they are ahead of time. Sorting by arbitrary fields means copypasting lots of boilerplate with small nuanced changes.

Sorting with `tros`
-----

```
type Thing struct {
	Name           string
	Weight, Height int
	Awesome        bool
}

func sortThings(things []Thing) {
    fmt.Println("Sorted by name:")
    tros.Sort(things, "Name")
    fmt.Println(things)
    
    fmt.Println("Sorted by weight:")
    tros.Sort(things, "Weight")
    fmt.Println(things)
    
    // Sort in reverse order too!
    fmt.Println("Sorted by name (reversed):")
    tros.Sort(things, "-Name")
    fmt.Println(things)
    
    // ...and so on
}
```

With `tros`, you can easily sort slices of structs by the values of arbitrary fields described at runtime, leaving out lots of boring boilerplate.

Custom Sorting
-----
If you want to define your own custom sorting logic, you can have your struct fields implement the [`Lesser`](https://godoc.org/github.com/ImJasonH/tros#Lesser) interface, which `tros` will use to determine ordering.

Caveats
-----

* `tros` uses reflection, and as such is noticeably (~2.5x) slower for sorting than the idiomatic Go way. But for small-to-medium slices, the difference should be negligible. If performance is more important than convenience, use the standard [`sort`](https://godoc.org/sort) package


----------

License
-----

    Copyright 2014 Jason Hall

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.

