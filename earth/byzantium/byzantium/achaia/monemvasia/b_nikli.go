package monemvasia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼克拉斯NikliBarony struct {
	feud.BaseBarony
}

var BNikli尼克拉斯 feud.Barony = &尼克拉斯NikliBarony{}

func init() {
    f := BNikli尼克拉斯.(*尼克拉斯NikliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikli",
		TitleName: "尼克拉斯",
		TitleCode: "b_nikli",
	}
}
