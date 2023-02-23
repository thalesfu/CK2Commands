package tabriz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大不里士TabrizBarony struct {
	feud.BaseBarony
}

var BTabriz大不里士 feud.Barony = &大不里士TabrizBarony{}

func init() {
	f := BTabriz大不里士.(*大不里士TabrizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tabriz",
		TitleName: "大不里士",
		TitleCode: "b_tabriz",
	}
}
