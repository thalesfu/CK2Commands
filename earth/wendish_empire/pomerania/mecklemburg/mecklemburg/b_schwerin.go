package mecklemburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什未林SchwerinBarony struct {
	feud.BaseBarony
}

var BSchwerin什未林 feud.Barony = &什未林SchwerinBarony{}

func init() {
    f := BSchwerin什未林.(*什未林SchwerinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schwerin",
		TitleName: "什未林",
		TitleCode: "b_schwerin",
	}
}
