package geneve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧博讷AubonneBarony struct {
	feud.BaseBarony
}

var BAubonne欧博讷 feud.Barony = &欧博讷AubonneBarony{}

func init() {
	f := BAubonne欧博讷.(*欧博讷AubonneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aubonne",
		TitleName: "欧博讷",
		TitleCode: "b_aubonne",
	}
}
