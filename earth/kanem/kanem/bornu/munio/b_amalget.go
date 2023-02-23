package munio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿马尔盖特AmalgetBarony struct {
	feud.BaseBarony
}

var BAmalget阿马尔盖特 feud.Barony = &阿马尔盖特AmalgetBarony{}

func init() {
	f := BAmalget阿马尔盖特.(*阿马尔盖特AmalgetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amalget",
		TitleName: "阿马尔盖特",
		TitleCode: "b_amalget",
	}
}
