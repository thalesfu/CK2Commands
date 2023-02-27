package lesbos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普洛马里PlomariBarony struct {
	feud.BaseBarony
}

var BPlomari普洛马里 feud.Barony = &普洛马里PlomariBarony{}

func init() {
    f := BPlomari普洛马里.(*普洛马里PlomariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plomari",
		TitleName: "普洛马里",
		TitleCode: "b_plomari",
	}
}
