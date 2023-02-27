package sambia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柯尼斯堡KonigsbergBarony struct {
	feud.BaseBarony
}

var BKonigsberg柯尼斯堡 feud.Barony = &柯尼斯堡KonigsbergBarony{}

func init() {
    f := BKonigsberg柯尼斯堡.(*柯尼斯堡KonigsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konigsberg",
		TitleName: "柯尼斯堡",
		TitleCode: "b_konigsberg",
	}
}
