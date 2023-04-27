# Subject ID

Subject Identifiers for Security Event Tokens, implementing [draft-ietf-secevent-subject-identifiers](https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers-16).

This is a nearly feature-complete implementation of draft-16, including the given examples as unit tests.

The code follows draft-16, which may be the Editor's copy rather than the published draft.

## Usage

The library provides a simple api via the ID type as well as a set of constructors corresponding to each subject id format.

Given below is a simple example showing sample usage
through the constructor:

```go
package main

import (
    "github.com/bingxueshuang/gnap/subject"
)

func main() {
    // handle errors appropriately.
    isssub, _ := subject.NewIDIssSub("https://identity,example.org", "FNJ45HJ6")
    emailid, _ := subject.NewIDEmail("editor@example.org")
    userinfo, _ := subject.NewAliases([]subject.NoAlias{
        isssub.NoAlias(),
        emailid.NoAlias(),
    })
}
```

For more complicated use cases, directly use the ID type
or the NoAlias type.

Validator method is given so as to allow a means to check if the subject id fields are valid or not.

## Links

- [Subject Identifiers for Security Event Tokens draft](https://datatracker.ietf.org/doc/draft-ietf-secevent-subject-identifiers)
