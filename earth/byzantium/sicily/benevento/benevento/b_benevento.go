package benevento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝内文托BeneventoBarony struct {
	feud.BaseBarony
}

var BBenevento贝内文托 feud.Barony = &贝内文托BeneventoBarony{}

func init() {
	f := BBenevento贝内文托.(*贝内文托BeneventoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "benevento",
		TitleName: "贝内文托",
		TitleCode: "b_benevento",
	}
}
