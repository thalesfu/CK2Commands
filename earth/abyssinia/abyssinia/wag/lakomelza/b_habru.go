package lakomelza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈布茹HabruBarony struct {
	feud.BaseBarony
}

var BHabru哈布茹 feud.Barony = &哈布茹HabruBarony{}

func init() {
    f := BHabru哈布茹.(*哈布茹HabruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "habru",
		TitleName: "哈布茹",
		TitleCode: "b_habru",
	}
}
