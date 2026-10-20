// Package spec holds the Flipt-specific assets for the OpenFeature Provider TCK.
package spec

import (
	"embed"
	"io/fs"
)

//go:embed extensions/*.feature
var featureFiles embed.FS

// Extensions holds the Flipt extension scenarios. The filesystem is rooted at
// the feature files, so tck.WithFeatures(spec.Extensions) reads them directly.
var Extensions fs.FS = mustSub(featureFiles, "extensions")

func mustSub(fsys fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err)
	}
	return sub
}
