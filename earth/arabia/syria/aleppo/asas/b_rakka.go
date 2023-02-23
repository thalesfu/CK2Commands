package asas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 腊卡RakkaBarony struct {
	feud.BaseBarony
}

var BRakka腊卡 feud.Barony = &腊卡RakkaBarony{}

func init() {
	f := BRakka腊卡.(*腊卡RakkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rakka",
		TitleName: "腊卡",
		TitleCode: "b_rakka",
	}
}
