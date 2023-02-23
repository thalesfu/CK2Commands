package purang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多隆DulungBarony struct {
	feud.BaseBarony
}

var BDulung多隆 feud.Barony = &多隆DulungBarony{}

func init() {
	f := BDulung多隆.(*多隆DulungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dulung",
		TitleName: "多隆",
		TitleCode: "b_dulung",
	}
}
