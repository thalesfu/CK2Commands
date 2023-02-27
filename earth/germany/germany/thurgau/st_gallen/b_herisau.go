package st_gallen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑里绍HerisauBarony struct {
	feud.BaseBarony
}

var BHerisau黑里绍 feud.Barony = &黑里绍HerisauBarony{}

func init() {
    f := BHerisau黑里绍.(*黑里绍HerisauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "herisau",
		TitleName: "黑里绍",
		TitleCode: "b_herisau",
	}
}
