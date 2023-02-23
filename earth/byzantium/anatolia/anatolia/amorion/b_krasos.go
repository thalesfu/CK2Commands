package amorion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉索斯KrasosBarony struct {
	feud.BaseBarony
}

var BKrasos克拉索斯 feud.Barony = &克拉索斯KrasosBarony{}

func init() {
	f := BKrasos克拉索斯.(*克拉索斯KrasosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krasos",
		TitleName: "克拉索斯",
		TitleCode: "b_krasos",
	}
}
