package kastoria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯拉拉KailaraBarony struct {
	feud.BaseBarony
}

var BKailara凯拉拉 feud.Barony = &凯拉拉KailaraBarony{}

func init() {
	f := BKailara凯拉拉.(*凯拉拉KailaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kailara",
		TitleName: "凯拉拉",
		TitleCode: "b_kailara",
	}
}
