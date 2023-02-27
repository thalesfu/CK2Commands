package kozelsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索先斯基SosenskyBarony struct {
	feud.BaseBarony
}

var BSosensky索先斯基 feud.Barony = &索先斯基SosenskyBarony{}

func init() {
    f := BSosensky索先斯基.(*索先斯基SosenskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sosensky",
		TitleName: "索先斯基",
		TitleCode: "b_sosensky",
	}
}
