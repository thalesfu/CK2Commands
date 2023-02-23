package sancerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科讷CosneBarony struct {
	feud.BaseBarony
}

var BCosne科讷 feud.Barony = &科讷CosneBarony{}

func init() {
	f := BCosne科讷.(*科讷CosneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cosne",
		TitleName: "科讷",
		TitleCode: "b_cosne",
	}
}
