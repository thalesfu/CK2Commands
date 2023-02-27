package zawila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大佩特拉斯Petras_maiorBarony struct {
	feud.BaseBarony
}

var BPetras_maior大佩特拉斯 feud.Barony = &大佩特拉斯Petras_maiorBarony{}

func init() {
    f := BPetras_maior大佩特拉斯.(*大佩特拉斯Petras_maiorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petras_maior",
		TitleName: "大佩特拉斯",
		TitleCode: "b_petras_maior",
	}
}
