package adapter

import (
	"fmt"

	"github.com/xinnjie/confgen_by_parser/confgen"
)

type Cell interface {
	fmt.Stringer

	Pos() confgen.Position
}
