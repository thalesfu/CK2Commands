package antiocheia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈里姆HarencBarony struct {
	feud.BaseBarony
}

var BHarenc哈里姆 feud.Barony = &哈里姆HarencBarony{}

func init() {
	f := BHarenc哈里姆.(*哈里姆HarencBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harenc",
		TitleName: "哈里姆",
		TitleCode: "b_harenc",
	}
}
