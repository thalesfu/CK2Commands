package hardsyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱姆维LemvigBarony struct {
	feud.BaseBarony
}

var BLemvig莱姆维 feud.Barony = &莱姆维LemvigBarony{}

func init() {
	f := BLemvig莱姆维.(*莱姆维LemvigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lemvig",
		TitleName: "莱姆维",
		TitleCode: "b_lemvig",
	}
}
