package perm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔登CherdynBarony struct {
	feud.BaseBarony
}

var BCherdyn切尔登 feud.Barony = &切尔登CherdynBarony{}

func init() {
    f := BCherdyn切尔登.(*切尔登CherdynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherdyn",
		TitleName: "切尔登",
		TitleCode: "b_cherdyn",
	}
}
