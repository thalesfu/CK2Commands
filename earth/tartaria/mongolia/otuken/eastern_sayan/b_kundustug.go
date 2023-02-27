package eastern_sayan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昆杜斯图格KundustugBarony struct {
	feud.BaseBarony
}

var BKundustug昆杜斯图格 feud.Barony = &昆杜斯图格KundustugBarony{}

func init() {
    f := BKundustug昆杜斯图格.(*昆杜斯图格KundustugBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kundustug",
		TitleName: "昆杜斯图格",
		TitleCode: "b_kundustug",
	}
}
