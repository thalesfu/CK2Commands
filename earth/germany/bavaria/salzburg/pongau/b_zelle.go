package pongau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采尔ZelleBarony struct {
	feud.BaseBarony
}

var BZelle采尔 feud.Barony = &采尔ZelleBarony{}

func init() {
    f := BZelle采尔.(*采尔ZelleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zelle",
		TitleName: "采尔",
		TitleCode: "b_zelle",
	}
}
