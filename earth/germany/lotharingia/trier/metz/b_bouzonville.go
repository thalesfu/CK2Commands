package metz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布宗维勒BouzonvilleBarony struct {
	feud.BaseBarony
}

var BBouzonville布宗维勒 feud.Barony = &布宗维勒BouzonvilleBarony{}

func init() {
    f := BBouzonville布宗维勒.(*布宗维勒BouzonvilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bouzonville",
		TitleName: "布宗维勒",
		TitleCode: "b_bouzonville",
	}
}
