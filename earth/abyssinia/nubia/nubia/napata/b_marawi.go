package napata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麦罗维MarawiBarony struct {
	feud.BaseBarony
}

var BMarawi麦罗维 feud.Barony = &麦罗维MarawiBarony{}

func init() {
	f := BMarawi麦罗维.(*麦罗维MarawiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marawi",
		TitleName: "麦罗维",
		TitleCode: "b_marawi",
	}
}
