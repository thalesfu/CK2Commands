package vinnytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯沙德BershadBarony struct {
	feud.BaseBarony
}

var BBershad伯沙德 feud.Barony = &伯沙德BershadBarony{}

func init() {
    f := BBershad伯沙德.(*伯沙德BershadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bershad",
		TitleName: "伯沙德",
		TitleCode: "b_bershad",
	}
}
