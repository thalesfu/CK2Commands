package mathura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奎师那诞生地KrishnajanmabhoomiBarony struct {
	feud.BaseBarony
}

var BKrishnajanmabhoomi奎师那诞生地 feud.Barony = &奎师那诞生地KrishnajanmabhoomiBarony{}

func init() {
	f := BKrishnajanmabhoomi奎师那诞生地.(*奎师那诞生地KrishnajanmabhoomiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krishnajanmabhoomi",
		TitleName: "奎师那诞生地",
		TitleCode: "b_krishnajanmabhoomi",
	}
}
