package consenza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔真塔诺ArgentanoBarony struct {
	feud.BaseBarony
}

var BArgentano阿尔真塔诺 feud.Barony = &阿尔真塔诺ArgentanoBarony{}

func init() {
	f := BArgentano阿尔真塔诺.(*阿尔真塔诺ArgentanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "argentano",
		TitleName: "阿尔真塔诺",
		TitleCode: "b_argentano",
	}
}
