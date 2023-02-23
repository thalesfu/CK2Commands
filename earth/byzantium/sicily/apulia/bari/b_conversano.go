package bari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔韦尔萨诺ConversanoBarony struct {
	feud.BaseBarony
}

var BConversano孔韦尔萨诺 feud.Barony = &孔韦尔萨诺ConversanoBarony{}

func init() {
	f := BConversano孔韦尔萨诺.(*孔韦尔萨诺ConversanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "conversano",
		TitleName: "孔韦尔萨诺",
		TitleCode: "b_conversano",
	}
}
