package kermanshah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡拉万HulwanBarony struct {
	feud.BaseBarony
}

var BHulwan胡拉万 feud.Barony = &胡拉万HulwanBarony{}

func init() {
    f := BHulwan胡拉万.(*胡拉万HulwanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hulwan",
		TitleName: "胡拉万",
		TitleCode: "b_hulwan",
	}
}
