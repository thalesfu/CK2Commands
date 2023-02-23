package kaliopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加里波利GallipoliBarony struct {
	feud.BaseBarony
}

var BGallipoli加里波利 feud.Barony = &加里波利GallipoliBarony{}

func init() {
	f := BGallipoli加里波利.(*加里波利GallipoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gallipoli",
		TitleName: "加里波利",
		TitleCode: "b_gallipoli",
	}
}
