package soso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨马SamaBarony struct {
	feud.BaseBarony
}

var BSama萨马 feud.Barony = &萨马SamaBarony{}

func init() {
    f := BSama萨马.(*萨马SamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sama",
		TitleName: "萨马",
		TitleCode: "b_sama",
	}
}
