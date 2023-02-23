package qwivir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达姆甘DamqanBarony struct {
	feud.BaseBarony
}

var BDamqan达姆甘 feud.Barony = &达姆甘DamqanBarony{}

func init() {
	f := BDamqan达姆甘.(*达姆甘DamqanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damqan",
		TitleName: "达姆甘",
		TitleCode: "b_damqan",
	}
}
