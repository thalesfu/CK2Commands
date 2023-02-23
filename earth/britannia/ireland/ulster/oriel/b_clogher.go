package oriel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克洛赫ClogherBarony struct {
	feud.BaseBarony
}

var BClogher克洛赫 feud.Barony = &克洛赫ClogherBarony{}

func init() {
	f := BClogher克洛赫.(*克洛赫ClogherBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clogher",
		TitleName: "克洛赫",
		TitleCode: "b_clogher",
	}
}
