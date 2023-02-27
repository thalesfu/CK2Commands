package sarangpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩梨耶佛俱BhoilymbongBarony struct {
	feud.BaseBarony
}

var BBhoilymbong菩梨耶佛俱 feud.Barony = &菩梨耶佛俱BhoilymbongBarony{}

func init() {
    f := BBhoilymbong菩梨耶佛俱.(*菩梨耶佛俱BhoilymbongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhoilymbong",
		TitleName: "菩梨耶佛俱",
		TitleCode: "b_bhoilymbong",
	}
}
