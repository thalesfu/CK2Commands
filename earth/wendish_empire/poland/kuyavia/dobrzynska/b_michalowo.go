package dobrzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米哈沃沃MichalowoBarony struct {
	feud.BaseBarony
}

var BMichalowo米哈沃沃 feud.Barony = &米哈沃沃MichalowoBarony{}

func init() {
    f := BMichalowo米哈沃沃.(*米哈沃沃MichalowoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "michalowo",
		TitleName: "米哈沃沃",
		TitleCode: "b_michalowo",
	}
}
