package djimi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马奥MaoBarony struct {
	feud.BaseBarony
}

var BMao马奥 feud.Barony = &马奥MaoBarony{}

func init() {
	f := BMao马奥.(*马奥MaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mao",
		TitleName: "马奥",
		TitleCode: "b_mao",
	}
}
