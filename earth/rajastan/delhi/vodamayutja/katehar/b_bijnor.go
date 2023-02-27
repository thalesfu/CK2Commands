package katehar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗杰诺尔BijnorBarony struct {
	feud.BaseBarony
}

var BBijnor毗杰诺尔 feud.Barony = &毗杰诺尔BijnorBarony{}

func init() {
    f := BBijnor毗杰诺尔.(*毗杰诺尔BijnorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bijnor",
		TitleName: "毗杰诺尔",
		TitleCode: "b_bijnor",
	}
}
