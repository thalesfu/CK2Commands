package uwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜尔AldurBarony struct {
	feud.BaseBarony
}

var BAldur杜尔 feud.Barony = &杜尔AldurBarony{}

func init() {
    f := BAldur杜尔.(*杜尔AldurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aldur",
		TitleName: "杜尔",
		TitleCode: "b_aldur",
	}
}
