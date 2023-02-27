package pangong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克赛钦AksayqinBarony struct {
	feud.BaseBarony
}

var BAksayqin阿克赛钦 feud.Barony = &阿克赛钦AksayqinBarony{}

func init() {
    f := BAksayqin阿克赛钦.(*阿克赛钦AksayqinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aksayqin",
		TitleName: "阿克赛钦",
		TitleCode: "b_aksayqin",
	}
}
