package nizhny_novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尼亚吉尼诺KnyagininoBarony struct {
	feud.BaseBarony
}

var BKnyaginino克尼亚吉尼诺 feud.Barony = &克尼亚吉尼诺KnyagininoBarony{}

func init() {
    f := BKnyaginino克尼亚吉尼诺.(*克尼亚吉尼诺KnyagininoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knyaginino",
		TitleName: "克尼亚吉尼诺",
		TitleCode: "b_knyaginino",
	}
}
