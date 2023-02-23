package tortosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉克利MaracleaBarony struct {
	feud.BaseBarony
}

var BMaraclea马拉克利 feud.Barony = &马拉克利MaracleaBarony{}

func init() {
	f := BMaraclea马拉克利.(*马拉克利MaracleaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maraclea",
		TitleName: "马拉克利",
		TitleCode: "b_maraclea",
	}
}
