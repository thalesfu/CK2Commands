package manyakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑翁耆SaongiBarony struct {
	feud.BaseBarony
}

var BSaongi娑翁耆 feud.Barony = &娑翁耆SaongiBarony{}

func init() {
	f := BSaongi娑翁耆.(*娑翁耆SaongiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saongi",
		TitleName: "娑翁耆",
		TitleCode: "b_saongi",
	}
}
