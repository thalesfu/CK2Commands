package ross

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰恩TainBarony struct {
	feud.BaseBarony
}

var BTain泰恩 feud.Barony = &泰恩TainBarony{}

func init() {
    f := BTain泰恩.(*泰恩TainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tain",
		TitleName: "泰恩",
		TitleCode: "b_tain",
	}
}
