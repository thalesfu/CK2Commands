package tyrone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德里DerryBarony struct {
	feud.BaseBarony
}

var BDerry德里 feud.Barony = &德里DerryBarony{}

func init() {
	f := BDerry德里.(*德里DerryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derry",
		TitleName: "德里",
		TitleCode: "b_derry",
	}
}
