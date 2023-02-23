package oxford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾尔斯伯里AylesburyBarony struct {
	feud.BaseBarony
}

var BAylesbury艾尔斯伯里 feud.Barony = &艾尔斯伯里AylesburyBarony{}

func init() {
	f := BAylesbury艾尔斯伯里.(*艾尔斯伯里AylesburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aylesbury",
		TitleName: "艾尔斯伯里",
		TitleCode: "b_aylesbury",
	}
}
