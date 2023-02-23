package builth

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伯雷杜AberedwBarony struct {
	feud.BaseBarony
}

var BAberedw阿伯雷杜 feud.Barony = &阿伯雷杜AberedwBarony{}

func init() {
	f := BAberedw阿伯雷杜.(*阿伯雷杜AberedwBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aberedw",
		TitleName: "阿伯雷杜",
		TitleCode: "b_aberedw",
	}
}
