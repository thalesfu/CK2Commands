package bourbon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波旁BourbonBarony struct {
	feud.BaseBarony
}

var BBourbon波旁 feud.Barony = &波旁BourbonBarony{}

func init() {
	f := BBourbon波旁.(*波旁BourbonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bourbon",
		TitleName: "波旁",
		TitleCode: "b_bourbon",
	}
}
