package hamadan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁德阿瓦尔RoudavarBarony struct {
	feud.BaseBarony
}

var BRoudavar鲁德阿瓦尔 feud.Barony = &鲁德阿瓦尔RoudavarBarony{}

func init() {
	f := BRoudavar鲁德阿瓦尔.(*鲁德阿瓦尔RoudavarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roudavar",
		TitleName: "鲁德阿瓦尔",
		TitleCode: "b_roudavar",
	}
}
