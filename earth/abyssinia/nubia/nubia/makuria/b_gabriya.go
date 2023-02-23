package makuria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽比亚GabriyaBarony struct {
	feud.BaseBarony
}

var BGabriya伽比亚 feud.Barony = &伽比亚GabriyaBarony{}

func init() {
	f := BGabriya伽比亚.(*伽比亚GabriyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gabriya",
		TitleName: "伽比亚",
		TitleCode: "b_gabriya",
	}
}
