package ulster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 顿塞沃里克DunseverickBarony struct {
	feud.BaseBarony
}

var BDunseverick顿塞沃里克 feud.Barony = &顿塞沃里克DunseverickBarony{}

func init() {
	f := BDunseverick顿塞沃里克.(*顿塞沃里克DunseverickBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunseverick",
		TitleName: "顿塞沃里克",
		TitleCode: "b_dunseverick",
	}
}
