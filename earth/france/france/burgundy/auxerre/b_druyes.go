package auxerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德吕耶DruyesBarony struct {
	feud.BaseBarony
}

var BDruyes德吕耶 feud.Barony = &德吕耶DruyesBarony{}

func init() {
	f := BDruyes德吕耶.(*德吕耶DruyesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "druyes",
		TitleName: "德吕耶",
		TitleCode: "b_druyes",
	}
}
