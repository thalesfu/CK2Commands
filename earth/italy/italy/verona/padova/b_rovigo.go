package padova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗维戈RovigoBarony struct {
	feud.BaseBarony
}

var BRovigo罗维戈 feud.Barony = &罗维戈RovigoBarony{}

func init() {
	f := BRovigo罗维戈.(*罗维戈RovigoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rovigo",
		TitleName: "罗维戈",
		TitleCode: "b_rovigo",
	}
}
