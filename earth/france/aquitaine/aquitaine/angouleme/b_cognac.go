package angouleme

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 干邑CognacBarony struct {
	feud.BaseBarony
}

var BCognac干邑 feud.Barony = &干邑CognacBarony{}

func init() {
	f := BCognac干邑.(*干邑CognacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cognac",
		TitleName: "干邑",
		TitleCode: "b_cognac",
	}
}
