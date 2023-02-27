package carcassonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔卡松CarcassonneBarony struct {
	feud.BaseBarony
}

var BCarcassonne卡尔卡松 feud.Barony = &卡尔卡松CarcassonneBarony{}

func init() {
    f := BCarcassonne卡尔卡松.(*卡尔卡松CarcassonneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carcassonne",
		TitleName: "卡尔卡松",
		TitleCode: "b_carcassonne",
	}
}
