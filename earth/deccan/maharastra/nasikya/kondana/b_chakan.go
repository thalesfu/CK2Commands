package kondana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查甘ChakanBarony struct {
	feud.BaseBarony
}

var BChakan查甘 feud.Barony = &查甘ChakanBarony{}

func init() {
	f := BChakan查甘.(*查甘ChakanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chakan",
		TitleName: "查甘",
		TitleCode: "b_chakan",
	}
}
