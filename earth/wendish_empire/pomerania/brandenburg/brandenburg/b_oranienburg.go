package brandenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥拉宁堡OranienburgBarony struct {
	feud.BaseBarony
}

var BOranienburg奥拉宁堡 feud.Barony = &奥拉宁堡OranienburgBarony{}

func init() {
    f := BOranienburg奥拉宁堡.(*奥拉宁堡OranienburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oranienburg",
		TitleName: "奥拉宁堡",
		TitleCode: "b_oranienburg",
	}
}
