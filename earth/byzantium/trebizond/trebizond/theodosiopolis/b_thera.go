package theodosiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷TheraBarony struct {
	feud.BaseBarony
}

var BThera特雷 feud.Barony = &特雷TheraBarony{}

func init() {
    f := BThera特雷.(*特雷TheraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thera",
		TitleName: "特雷",
		TitleCode: "b_thera",
	}
}
