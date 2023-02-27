package magadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 华氏城PataliputraBarony struct {
	feud.BaseBarony
}

var BPataliputra华氏城 feud.Barony = &华氏城PataliputraBarony{}

func init() {
    f := BPataliputra华氏城.(*华氏城PataliputraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pataliputra",
		TitleName: "华氏城",
		TitleCode: "b_pataliputra",
	}
}
