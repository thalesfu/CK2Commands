package hendjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥米迪耶OmidiyehBarony struct {
	feud.BaseBarony
}

var BOmidiyeh奥米迪耶 feud.Barony = &奥米迪耶OmidiyehBarony{}

func init() {
	f := BOmidiyeh奥米迪耶.(*奥米迪耶OmidiyehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "omidiyeh",
		TitleName: "奥米迪耶",
		TitleCode: "b_omidiyeh",
	}
}
