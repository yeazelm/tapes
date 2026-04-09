package inmemory_test

import (
	"github.com/papercomputeco/tapes/pkg/storage"
	"github.com/papercomputeco/tapes/pkg/storage/inmemory"
	"github.com/papercomputeco/tapes/pkg/storage/storagetest"
)

var _ = storagetest.RunListSessionsSpecs("inmemory", func() storage.Driver {
	return inmemory.NewDriver()
})

var _ = storagetest.RunAncestryChainBasicSpecs("inmemory", func() storage.Driver {
	return inmemory.NewDriver()
})

var _ = storagetest.RunAncestryChainDanglingSpecs("inmemory", func() storage.Driver {
	return inmemory.NewDriver()
})
