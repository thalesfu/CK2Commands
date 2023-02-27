package belgrade

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁德尼克RudnikBarony struct {
	feud.BaseBarony
}

var BRudnik鲁德尼克 feud.Barony = &鲁德尼克RudnikBarony{}

func init() {
    f := BRudnik鲁德尼克.(*鲁德尼克RudnikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rudnik",
		TitleName: "鲁德尼克",
		TitleCode: "b_rudnik",
	}
}
