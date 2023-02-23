package vitebsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利奥兹诺LioznaBarony struct {
	feud.BaseBarony
}

var BLiozna利奥兹诺 feud.Barony = &利奥兹诺LioznaBarony{}

func init() {
	f := BLiozna利奥兹诺.(*利奥兹诺LioznaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liozna",
		TitleName: "利奥兹诺",
		TitleCode: "b_liozna",
	}
}
