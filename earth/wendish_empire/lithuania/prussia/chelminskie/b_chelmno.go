package chelminskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海乌姆诺ChelmnoBarony struct {
	feud.BaseBarony
}

var BChelmno海乌姆诺 feud.Barony = &海乌姆诺ChelmnoBarony{}

func init() {
    f := BChelmno海乌姆诺.(*海乌姆诺ChelmnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chelmno",
		TitleName: "海乌姆诺",
		TitleCode: "b_chelmno",
	}
}
