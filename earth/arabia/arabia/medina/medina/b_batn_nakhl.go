package medina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴特恩奈赫勒Batn_nakhlBarony struct {
	feud.BaseBarony
}

var BBatn_nakhl巴特恩奈赫勒 feud.Barony = &巴特恩奈赫勒Batn_nakhlBarony{}

func init() {
    f := BBatn_nakhl巴特恩奈赫勒.(*巴特恩奈赫勒Batn_nakhlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "batn_nakhl",
		TitleName: "巴特恩奈赫勒",
		TitleCode: "b_batn_nakhl",
	}
}
