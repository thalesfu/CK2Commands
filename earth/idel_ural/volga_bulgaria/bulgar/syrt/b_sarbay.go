package syrt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔拜SarbayBarony struct {
	feud.BaseBarony
}

var BSarbay萨尔拜 feud.Barony = &萨尔拜SarbayBarony{}

func init() {
    f := BSarbay萨尔拜.(*萨尔拜SarbayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarbay",
		TitleName: "萨尔拜",
		TitleCode: "b_sarbay",
	}
}
