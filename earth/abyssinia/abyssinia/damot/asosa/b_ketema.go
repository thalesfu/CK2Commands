package asosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科特马KetemaBarony struct {
	feud.BaseBarony
}

var BKetema科特马 feud.Barony = &科特马KetemaBarony{}

func init() {
	f := BKetema科特马.(*科特马KetemaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ketema",
		TitleName: "科特马",
		TitleCode: "b_ketema",
	}
}
