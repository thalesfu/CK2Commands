package tobruk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿雷马AcromaBarony struct {
	feud.BaseBarony
}

var BAcroma阿雷马 feud.Barony = &阿雷马AcromaBarony{}

func init() {
    f := BAcroma阿雷马.(*阿雷马AcromaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "acroma",
		TitleName: "阿雷马",
		TitleCode: "b_acroma",
	}
}
