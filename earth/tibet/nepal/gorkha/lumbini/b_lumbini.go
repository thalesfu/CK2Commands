package lumbini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 岚毗尼LumbiniBarony struct {
	feud.BaseBarony
}

var BLumbini岚毗尼 feud.Barony = &岚毗尼LumbiniBarony{}

func init() {
    f := BLumbini岚毗尼.(*岚毗尼LumbiniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lumbini",
		TitleName: "岚毗尼",
		TitleCode: "b_lumbini",
	}
}
