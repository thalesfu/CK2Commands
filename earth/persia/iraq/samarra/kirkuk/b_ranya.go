package kirkuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉尼耶RanyaBarony struct {
	feud.BaseBarony
}

var BRanya拉尼耶 feud.Barony = &拉尼耶RanyaBarony{}

func init() {
    f := BRanya拉尼耶.(*拉尼耶RanyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ranya",
		TitleName: "拉尼耶",
		TitleCode: "b_ranya",
	}
}
