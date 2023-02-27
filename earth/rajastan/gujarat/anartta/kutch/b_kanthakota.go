package kutch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 建多拘吒KanthakotaBarony struct {
	feud.BaseBarony
}

var BKanthakota建多拘吒 feud.Barony = &建多拘吒KanthakotaBarony{}

func init() {
    f := BKanthakota建多拘吒.(*建多拘吒KanthakotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanthakota",
		TitleName: "建多拘吒",
		TitleCode: "b_kanthakota",
	}
}
