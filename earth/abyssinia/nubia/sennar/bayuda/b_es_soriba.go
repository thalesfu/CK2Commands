package bayuda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索里巴Es_soribaBarony struct {
	feud.BaseBarony
}

var BEs_soriba索里巴 feud.Barony = &索里巴Es_soribaBarony{}

func init() {
    f := BEs_soriba索里巴.(*索里巴Es_soribaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "es_soriba",
		TitleName: "索里巴",
		TitleCode: "b_es_soriba",
	}
}
