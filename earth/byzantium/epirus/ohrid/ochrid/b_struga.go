package ochrid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特鲁加StrugaBarony struct {
	feud.BaseBarony
}

var BStruga斯特鲁加 feud.Barony = &斯特鲁加StrugaBarony{}

func init() {
	f := BStruga斯特鲁加.(*斯特鲁加StrugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "struga",
		TitleName: "斯特鲁加",
		TitleCode: "b_struga",
	}
}
