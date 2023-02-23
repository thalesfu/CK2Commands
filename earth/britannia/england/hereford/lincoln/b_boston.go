package lincoln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波士顿BostonBarony struct {
	feud.BaseBarony
}

var BBoston波士顿 feud.Barony = &波士顿BostonBarony{}

func init() {
	f := BBoston波士顿.(*波士顿BostonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boston",
		TitleName: "波士顿",
		TitleCode: "b_boston",
	}
}
