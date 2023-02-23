package nantes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 栋日DongesBarony struct {
	feud.BaseBarony
}

var BDonges栋日 feud.Barony = &栋日DongesBarony{}

func init() {
	f := BDonges栋日.(*栋日DongesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "donges",
		TitleName: "栋日",
		TitleCode: "b_donges",
	}
}
