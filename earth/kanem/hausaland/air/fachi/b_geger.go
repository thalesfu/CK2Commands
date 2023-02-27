package fachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖格尔GegerBarony struct {
	feud.BaseBarony
}

var BGeger盖格尔 feud.Barony = &盖格尔GegerBarony{}

func init() {
    f := BGeger盖格尔.(*盖格尔GegerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geger",
		TitleName: "盖格尔",
		TitleCode: "b_geger",
	}
}
