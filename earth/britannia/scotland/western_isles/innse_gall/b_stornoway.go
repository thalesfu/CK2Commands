package innse_gall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯托诺韦StornowayBarony struct {
	feud.BaseBarony
}

var BStornoway斯托诺韦 feud.Barony = &斯托诺韦StornowayBarony{}

func init() {
    f := BStornoway斯托诺韦.(*斯托诺韦StornowayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stornoway",
		TitleName: "斯托诺韦",
		TitleCode: "b_stornoway",
	}
}
