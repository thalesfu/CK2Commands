package biskra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马格拉MagraBarony struct {
	feud.BaseBarony
}

var BMagra马格拉 feud.Barony = &马格拉MagraBarony{}

func init() {
	f := BMagra马格拉.(*马格拉MagraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "magra",
		TitleName: "马格拉",
		TitleCode: "b_magra",
	}
}
