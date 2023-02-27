package tyrus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米格德尔MegedelBarony struct {
	feud.BaseBarony
}

var BMegedel米格德尔 feud.Barony = &米格德尔MegedelBarony{}

func init() {
    f := BMegedel米格德尔.(*米格德尔MegedelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "megedel",
		TitleName: "米格德尔",
		TitleCode: "b_megedel",
	}
}
