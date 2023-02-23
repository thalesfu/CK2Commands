package badghis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普尚格PushangBarony struct {
	feud.BaseBarony
}

var BPushang普尚格 feud.Barony = &普尚格PushangBarony{}

func init() {
	f := BPushang普尚格.(*普尚格PushangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pushang",
		TitleName: "普尚格",
		TitleCode: "b_pushang",
	}
}
