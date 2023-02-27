package al_amarah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖尔拜亚ZurbatiyahBarony struct {
	feud.BaseBarony
}

var BZurbatiyah祖尔拜亚 feud.Barony = &祖尔拜亚ZurbatiyahBarony{}

func init() {
    f := BZurbatiyah祖尔拜亚.(*祖尔拜亚ZurbatiyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zurbatiyah",
		TitleName: "祖尔拜亚",
		TitleCode: "b_zurbatiyah",
	}
}
