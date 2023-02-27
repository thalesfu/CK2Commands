package oleshye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克里维里赫Kryvyi_rihBarony struct {
	feud.BaseBarony
}

var BKryvyi_rih克里维里赫 feud.Barony = &克里维里赫Kryvyi_rihBarony{}

func init() {
    f := BKryvyi_rih克里维里赫.(*克里维里赫Kryvyi_rihBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kryvyi_rih",
		TitleName: "克里维里赫",
		TitleCode: "b_kryvyi_rih",
	}
}
