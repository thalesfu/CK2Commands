package yeruslan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 通格什TungyshBarony struct {
	feud.BaseBarony
}

var BTungysh通格什 feud.Barony = &通格什TungyshBarony{}

func init() {
    f := BTungysh通格什.(*通格什TungyshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tungysh",
		TitleName: "通格什",
		TitleCode: "b_tungysh",
	}
}
