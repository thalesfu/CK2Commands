package malta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆迪纳MdinaBarony struct {
	feud.BaseBarony
}

var BMdina姆迪纳 feud.Barony = &姆迪纳MdinaBarony{}

func init() {
	f := BMdina姆迪纳.(*姆迪纳MdinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mdina",
		TitleName: "姆迪纳",
		TitleCode: "b_mdina",
	}
}
