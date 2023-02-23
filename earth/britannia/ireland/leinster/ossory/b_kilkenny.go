package ossory

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔肯尼KilkennyBarony struct {
	feud.BaseBarony
}

var BKilkenny基尔肯尼 feud.Barony = &基尔肯尼KilkennyBarony{}

func init() {
	f := BKilkenny基尔肯尼.(*基尔肯尼KilkennyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kilkenny",
		TitleName: "基尔肯尼",
		TitleCode: "b_kilkenny",
	}
}
