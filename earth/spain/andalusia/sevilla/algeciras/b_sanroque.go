package algeciras

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣罗克SanroqueBarony struct {
	feud.BaseBarony
}

var BSanroque圣罗克 feud.Barony = &圣罗克SanroqueBarony{}

func init() {
	f := BSanroque圣罗克.(*圣罗克SanroqueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanroque",
		TitleName: "圣罗克",
		TitleCode: "b_sanroque",
	}
}
