package kasmira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿摩丽湿伐罗AmaresvaraBarony struct {
	feud.BaseBarony
}

var BAmaresvara阿摩丽湿伐罗 feud.Barony = &阿摩丽湿伐罗AmaresvaraBarony{}

func init() {
	f := BAmaresvara阿摩丽湿伐罗.(*阿摩丽湿伐罗AmaresvaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amaresvara",
		TitleName: "阿摩丽湿伐罗",
		TitleCode: "b_amaresvara",
	}
}
