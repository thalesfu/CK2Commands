package lutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒姆西克ShumskBarony struct {
	feud.BaseBarony
}

var BShumsk舒姆西克 feud.Barony = &舒姆西克ShumskBarony{}

func init() {
    f := BShumsk舒姆西克.(*舒姆西克ShumskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shumsk",
		TitleName: "舒姆西克",
		TitleCode: "b_shumsk",
	}
}
