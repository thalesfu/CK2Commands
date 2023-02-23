package nishapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赫拉马德AkhlamadBarony struct {
	feud.BaseBarony
}

var BAkhlamad阿赫拉马德 feud.Barony = &阿赫拉马德AkhlamadBarony{}

func init() {
	f := BAkhlamad阿赫拉马德.(*阿赫拉马德AkhlamadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akhlamad",
		TitleName: "阿赫拉马德",
		TitleCode: "b_akhlamad",
	}
}
