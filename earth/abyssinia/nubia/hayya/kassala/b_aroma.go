package kassala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿罗马AromaBarony struct {
	feud.BaseBarony
}

var BAroma阿罗马 feud.Barony = &阿罗马AromaBarony{}

func init() {
	f := BAroma阿罗马.(*阿罗马AromaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aroma",
		TitleName: "阿罗马",
		TitleCode: "b_aroma",
	}
}
