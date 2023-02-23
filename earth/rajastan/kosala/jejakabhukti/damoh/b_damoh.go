package damoh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佗牟DamohBarony struct {
	feud.BaseBarony
}

var BDamoh佗牟 feud.Barony = &佗牟DamohBarony{}

func init() {
	f := BDamoh佗牟.(*佗牟DamohBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damoh",
		TitleName: "佗牟",
		TitleCode: "b_damoh",
	}
}
