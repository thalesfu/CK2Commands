package lattalura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔BirBarony struct {
	feud.BaseBarony
}

var BBir比尔 feud.Barony = &比尔BirBarony{}

func init() {
	f := BBir比尔.(*比尔BirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bir",
		TitleName: "比尔",
		TitleCode: "b_bir",
	}
}
