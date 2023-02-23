package castellon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里亚纳BurrianaBarony struct {
	feud.BaseBarony
}

var BBurriana布里亚纳 feud.Barony = &布里亚纳BurrianaBarony{}

func init() {
	f := BBurriana布里亚纳.(*布里亚纳BurrianaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burriana",
		TitleName: "布里亚纳",
		TitleCode: "b_burriana",
	}
}
