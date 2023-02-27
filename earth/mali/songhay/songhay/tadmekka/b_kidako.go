package tadmekka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基达科KidakoBarony struct {
	feud.BaseBarony
}

var BKidako基达科 feud.Barony = &基达科KidakoBarony{}

func init() {
    f := BKidako基达科.(*基达科KidakoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kidako",
		TitleName: "基达科",
		TitleCode: "b_kidako",
	}
}
