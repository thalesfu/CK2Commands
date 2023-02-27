package smolensk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉斯内KrasnyBarony struct {
	feud.BaseBarony
}

var BKrasny克拉斯内 feud.Barony = &克拉斯内KrasnyBarony{}

func init() {
    f := BKrasny克拉斯内.(*克拉斯内KrasnyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krasny",
		TitleName: "克拉斯内",
		TitleCode: "b_krasny",
	}
}
