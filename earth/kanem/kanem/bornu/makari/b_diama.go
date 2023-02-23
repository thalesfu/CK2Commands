package makari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪亚马DiamaBarony struct {
	feud.BaseBarony
}

var BDiama迪亚马 feud.Barony = &迪亚马DiamaBarony{}

func init() {
	f := BDiama迪亚马.(*迪亚马DiamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diama",
		TitleName: "迪亚马",
		TitleCode: "b_diama",
	}
}
