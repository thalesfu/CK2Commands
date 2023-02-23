package noli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺利NoliBarony struct {
	feud.BaseBarony
}

var BNoli诺利 feud.Barony = &诺利NoliBarony{}

func init() {
	f := BNoli诺利.(*诺利NoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "noli",
		TitleName: "诺利",
		TitleCode: "b_noli",
	}
}
