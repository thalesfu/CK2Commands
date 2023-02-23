package methone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫顿ModonBarony struct {
	feud.BaseBarony
}

var BModon莫顿 feud.Barony = &莫顿ModonBarony{}

func init() {
	f := BModon莫顿.(*莫顿ModonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "modon",
		TitleName: "莫顿",
		TitleCode: "b_modon",
	}
}
