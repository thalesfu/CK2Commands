package lyon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布兰达BrindasBarony struct {
	feud.BaseBarony
}

var BBrindas布兰达 feud.Barony = &布兰达BrindasBarony{}

func init() {
	f := BBrindas布兰达.(*布兰达BrindasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brindas",
		TitleName: "布兰达",
		TitleCode: "b_brindas",
	}
}
