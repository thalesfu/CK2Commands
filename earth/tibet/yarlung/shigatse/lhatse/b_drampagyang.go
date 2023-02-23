package lhatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 仲巴江DrampagyangBarony struct {
	feud.BaseBarony
}

var BDrampagyang仲巴江 feud.Barony = &仲巴江DrampagyangBarony{}

func init() {
	f := BDrampagyang仲巴江.(*仲巴江DrampagyangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drampagyang",
		TitleName: "仲巴江",
		TitleCode: "b_drampagyang",
	}
}
