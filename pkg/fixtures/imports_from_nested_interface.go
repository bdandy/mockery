package test

import (
	"github.com/bdandy/mockery/v2/pkg/fixtures/http"
)

type HasConflictingNestedImports interface {
	RequesterNS
	Z() http.MyStruct
}
