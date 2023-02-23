package chandax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛利亚MaliaBarony struct {
	feud.BaseBarony
}

var BMalia玛利亚 feud.Barony = &玛利亚MaliaBarony{}

func init() {
	f := BMalia玛利亚.(*玛利亚MaliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malia",
		TitleName: "玛利亚",
		TitleCode: "b_malia",
	}
}
