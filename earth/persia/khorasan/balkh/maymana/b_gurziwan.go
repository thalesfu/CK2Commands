package maymana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔齐万GurziwanBarony struct {
	feud.BaseBarony
}

var BGurziwan古尔齐万 feud.Barony = &古尔齐万GurziwanBarony{}

func init() {
    f := BGurziwan古尔齐万.(*古尔齐万GurziwanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurziwan",
		TitleName: "古尔齐万",
		TitleCode: "b_gurziwan",
	}
}
