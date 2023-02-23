package valabhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弹度迦DhandhukaBarony struct {
	feud.BaseBarony
}

var BDhandhuka弹度迦 feud.Barony = &弹度迦DhandhukaBarony{}

func init() {
	f := BDhandhuka弹度迦.(*弹度迦DhandhukaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhandhuka",
		TitleName: "弹度迦",
		TitleCode: "b_dhandhuka",
	}
}
