package mecklemburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖沃钦GevezinBarony struct {
	feud.BaseBarony
}

var BGevezin盖沃钦 feud.Barony = &盖沃钦GevezinBarony{}

func init() {
    f := BGevezin盖沃钦.(*盖沃钦GevezinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gevezin",
		TitleName: "盖沃钦",
		TitleCode: "b_gevezin",
	}
}
