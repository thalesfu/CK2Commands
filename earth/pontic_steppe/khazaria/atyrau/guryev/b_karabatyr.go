package guryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉巴特尔KarabatyrBarony struct {
	feud.BaseBarony
}

var BKarabatyr哈拉巴特尔 feud.Barony = &哈拉巴特尔KarabatyrBarony{}

func init() {
    f := BKarabatyr哈拉巴特尔.(*哈拉巴特尔KarabatyrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karabatyr",
		TitleName: "哈拉巴特尔",
		TitleCode: "b_karabatyr",
	}
}
