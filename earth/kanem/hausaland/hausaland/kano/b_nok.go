package kano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺克NokBarony struct {
	feud.BaseBarony
}

var BNok诺克 feud.Barony = &诺克NokBarony{}

func init() {
	f := BNok诺克.(*诺克NokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nok",
		TitleName: "诺克",
		TitleCode: "b_nok",
	}
}
