package jeddah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡费宰QunfudhahBarony struct {
	feud.BaseBarony
}

var BQunfudhah贡费宰 feud.Barony = &贡费宰QunfudhahBarony{}

func init() {
	f := BQunfudhah贡费宰.(*贡费宰QunfudhahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qunfudhah",
		TitleName: "贡费宰",
		TitleCode: "b_qunfudhah",
	}
}
