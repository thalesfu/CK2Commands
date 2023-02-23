package kunggar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塘加TanggyaBarony struct {
	feud.BaseBarony
}

var BTanggya塘加 feud.Barony = &塘加TanggyaBarony{}

func init() {
	f := BTanggya塘加.(*塘加TanggyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tanggya",
		TitleName: "塘加",
		TitleCode: "b_tanggya",
	}
}
