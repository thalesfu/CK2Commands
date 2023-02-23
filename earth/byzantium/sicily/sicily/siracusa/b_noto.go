package siracusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺托NotoBarony struct {
	feud.BaseBarony
}

var BNoto诺托 feud.Barony = &诺托NotoBarony{}

func init() {
	f := BNoto诺托.(*诺托NotoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "noto",
		TitleName: "诺托",
		TitleCode: "b_noto",
	}
}
