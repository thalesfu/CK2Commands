package tyrone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥马OmaghBarony struct {
	feud.BaseBarony
}

var BOmagh奥马 feud.Barony = &奥马OmaghBarony{}

func init() {
	f := BOmagh奥马.(*奥马OmaghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "omagh",
		TitleName: "奥马",
		TitleCode: "b_omagh",
	}
}
