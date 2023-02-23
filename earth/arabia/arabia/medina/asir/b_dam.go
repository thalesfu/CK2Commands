package asir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达姆DamBarony struct {
	feud.BaseBarony
}

var BDam达姆 feud.Barony = &达姆DamBarony{}

func init() {
	f := BDam达姆.(*达姆DamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dam",
		TitleName: "达姆",
		TitleCode: "b_dam",
	}
}
