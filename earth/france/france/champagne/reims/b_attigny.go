package reims

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿蒂尼AttignyBarony struct {
	feud.BaseBarony
}

var BAttigny阿蒂尼 feud.Barony = &阿蒂尼AttignyBarony{}

func init() {
	f := BAttigny阿蒂尼.(*阿蒂尼AttignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "attigny",
		TitleName: "阿蒂尼",
		TitleCode: "b_attigny",
	}
}
