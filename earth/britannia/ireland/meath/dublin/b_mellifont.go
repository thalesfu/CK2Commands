package dublin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅利丰特MellifontBarony struct {
	feud.BaseBarony
}

var BMellifont梅利丰特 feud.Barony = &梅利丰特MellifontBarony{}

func init() {
	f := BMellifont梅利丰特.(*梅利丰特MellifontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mellifont",
		TitleName: "梅利丰特",
		TitleCode: "b_mellifont",
	}
}
