package saintonge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙吉永MontguyonBarony struct {
	feud.BaseBarony
}

var BMontguyon蒙吉永 feud.Barony = &蒙吉永MontguyonBarony{}

func init() {
	f := BMontguyon蒙吉永.(*蒙吉永MontguyonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montguyon",
		TitleName: "蒙吉永",
		TitleCode: "b_montguyon",
	}
}
