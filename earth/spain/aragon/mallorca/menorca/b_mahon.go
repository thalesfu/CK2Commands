package menorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马奥MahonBarony struct {
	feud.BaseBarony
}

var BMahon马奥 feud.Barony = &马奥MahonBarony{}

func init() {
	f := BMahon马奥.(*马奥MahonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahon",
		TitleName: "马奥",
		TitleCode: "b_mahon",
	}
}
