package akroinon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤门尼亚EumeniaBarony struct {
	feud.BaseBarony
}

var BEumenia尤门尼亚 feud.Barony = &尤门尼亚EumeniaBarony{}

func init() {
	f := BEumenia尤门尼亚.(*尤门尼亚EumeniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eumenia",
		TitleName: "尤门尼亚",
		TitleCode: "b_eumenia",
	}
}
