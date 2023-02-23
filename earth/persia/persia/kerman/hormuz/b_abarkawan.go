package hormuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴尔卡万AbarkawanBarony struct {
	feud.BaseBarony
}

var BAbarkawan阿巴尔卡万 feud.Barony = &阿巴尔卡万AbarkawanBarony{}

func init() {
	f := BAbarkawan阿巴尔卡万.(*阿巴尔卡万AbarkawanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abarkawan",
		TitleName: "阿巴尔卡万",
		TitleCode: "b_abarkawan",
	}
}
