package capua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰阿诺TeanoBarony struct {
	feud.BaseBarony
}

var BTeano泰阿诺 feud.Barony = &泰阿诺TeanoBarony{}

func init() {
	f := BTeano泰阿诺.(*泰阿诺TeanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teano",
		TitleName: "泰阿诺",
		TitleCode: "b_teano",
	}
}
