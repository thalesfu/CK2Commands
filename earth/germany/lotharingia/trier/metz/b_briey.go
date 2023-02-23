package metz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里埃BrieyBarony struct {
	feud.BaseBarony
}

var BBriey布里埃 feud.Barony = &布里埃BrieyBarony{}

func init() {
	f := BBriey布里埃.(*布里埃BrieyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "briey",
		TitleName: "布里埃",
		TitleCode: "b_briey",
	}
}
