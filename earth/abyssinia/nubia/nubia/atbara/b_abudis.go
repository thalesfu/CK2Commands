package atbara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布迪斯AbudisBarony struct {
	feud.BaseBarony
}

var BAbudis阿布迪斯 feud.Barony = &阿布迪斯AbudisBarony{}

func init() {
	f := BAbudis阿布迪斯.(*阿布迪斯AbudisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abudis",
		TitleName: "阿布迪斯",
		TitleCode: "b_abudis",
	}
}
