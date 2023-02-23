package dulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 都兰DulanBarony struct {
	feud.BaseBarony
}

var BDulan都兰 feud.Barony = &都兰DulanBarony{}

func init() {
	f := BDulan都兰.(*都兰DulanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dulan",
		TitleName: "都兰",
		TitleCode: "b_dulan",
	}
}
