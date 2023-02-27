package leix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 芒特拉斯MountrathBarony struct {
	feud.BaseBarony
}

var BMountrath芒特拉斯 feud.Barony = &芒特拉斯MountrathBarony{}

func init() {
    f := BMountrath芒特拉斯.(*芒特拉斯MountrathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mountrath",
		TitleName: "芒特拉斯",
		TitleCode: "b_mountrath",
	}
}
