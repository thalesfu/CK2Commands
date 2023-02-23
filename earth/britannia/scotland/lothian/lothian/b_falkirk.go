package lothian

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福尔柯克FalkirkBarony struct {
	feud.BaseBarony
}

var BFalkirk福尔柯克 feud.Barony = &福尔柯克FalkirkBarony{}

func init() {
	f := BFalkirk福尔柯克.(*福尔柯克FalkirkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "falkirk",
		TitleName: "福尔柯克",
		TitleCode: "b_falkirk",
	}
}
