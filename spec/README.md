# spec

A separate Go module for the Flipt vendor scenarios.

The module path is `github.com/flipt-io/openfeature-tck/spec`. `Extensions`
embeds `extensions/*.feature` and is rooted at the feature files. A Go adoption
passes it to the suite:

```go
tck.WithFeatures(spec.Extensions)
```

`extensions/flipt-reason.feature` mirrors the canonical standard-reason rows
with the actual Flipt behavior. Update it by hand when the spec pin in the root
`go.mod` changes `reason.feature`.
