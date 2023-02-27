package middlesex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔西ChelseaBarony struct {
	feud.BaseBarony
}

var BChelsea切尔西 feud.Barony = &切尔西ChelseaBarony{}

func init() {
    f := BChelsea切尔西.(*切尔西ChelseaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chelsea",
		TitleName: "切尔西",
		TitleCode: "b_chelsea",
	}
}
