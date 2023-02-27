package malta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖里科MgarrBarony struct {
	feud.BaseBarony
}

var BMgarr祖里科 feud.Barony = &祖里科MgarrBarony{}

func init() {
    f := BMgarr祖里科.(*祖里科MgarrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mgarr",
		TitleName: "祖里科",
		TitleCode: "b_mgarr",
	}
}
