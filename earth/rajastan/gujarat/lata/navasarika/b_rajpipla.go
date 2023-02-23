package navasarika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗阇毕钵罗RajpiplaBarony struct {
	feud.BaseBarony
}

var BRajpipla罗阇毕钵罗 feud.Barony = &罗阇毕钵罗RajpiplaBarony{}

func init() {
	f := BRajpipla罗阇毕钵罗.(*罗阇毕钵罗RajpiplaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajpipla",
		TitleName: "罗阇毕钵罗",
		TitleCode: "b_rajpipla",
	}
}
