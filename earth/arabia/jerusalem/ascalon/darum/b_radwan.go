package darum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉德万RadwanBarony struct {
	feud.BaseBarony
}

var BRadwan拉德万 feud.Barony = &拉德万RadwanBarony{}

func init() {
    f := BRadwan拉德万.(*拉德万RadwanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radwan",
		TitleName: "拉德万",
		TitleCode: "b_radwan",
	}
}
