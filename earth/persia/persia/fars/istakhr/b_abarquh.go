package istakhr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴尔古AbarquhBarony struct {
	feud.BaseBarony
}

var BAbarquh阿巴尔古 feud.Barony = &阿巴尔古AbarquhBarony{}

func init() {
    f := BAbarquh阿巴尔古.(*阿巴尔古AbarquhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abarquh",
		TitleName: "阿巴尔古",
		TitleCode: "b_abarquh",
	}
}
