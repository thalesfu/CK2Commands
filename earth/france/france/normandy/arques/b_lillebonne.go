package arques

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利勒博讷LillebonneBarony struct {
	feud.BaseBarony
}

var BLillebonne利勒博讷 feud.Barony = &利勒博讷LillebonneBarony{}

func init() {
    f := BLillebonne利勒博讷.(*利勒博讷LillebonneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lillebonne",
		TitleName: "利勒博讷",
		TitleCode: "b_lillebonne",
	}
}
