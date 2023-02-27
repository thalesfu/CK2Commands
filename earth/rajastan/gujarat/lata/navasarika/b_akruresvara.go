package navasarika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拘卢丽湿伐罗AkruresvaraBarony struct {
	feud.BaseBarony
}

var BAkruresvara阿拘卢丽湿伐罗 feud.Barony = &阿拘卢丽湿伐罗AkruresvaraBarony{}

func init() {
    f := BAkruresvara阿拘卢丽湿伐罗.(*阿拘卢丽湿伐罗AkruresvaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akruresvara",
		TitleName: "阿拘卢丽湿伐罗",
		TitleCode: "b_akruresvara",
	}
}
