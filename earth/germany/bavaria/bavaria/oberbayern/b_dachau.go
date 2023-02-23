package oberbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达豪DachauBarony struct {
	feud.BaseBarony
}

var BDachau达豪 feud.Barony = &达豪DachauBarony{}

func init() {
	f := BDachau达豪.(*达豪DachauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dachau",
		TitleName: "达豪",
		TitleCode: "b_dachau",
	}
}
