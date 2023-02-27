package reni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗陀娑BidasarBarony struct {
	feud.BaseBarony
}

var BBidasar毗陀娑 feud.Barony = &毗陀娑BidasarBarony{}

func init() {
    f := BBidasar毗陀娑.(*毗陀娑BidasarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bidasar",
		TitleName: "毗陀娑",
		TitleCode: "b_bidasar",
	}
}
