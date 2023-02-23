package noli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨沃纳SavonaBarony struct {
	feud.BaseBarony
}

var BSavona萨沃纳 feud.Barony = &萨沃纳SavonaBarony{}

func init() {
	f := BSavona萨沃纳.(*萨沃纳SavonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "savona",
		TitleName: "萨沃纳",
		TitleCode: "b_savona",
	}
}
