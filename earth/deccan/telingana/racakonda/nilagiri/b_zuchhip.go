package nilagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叙侈波ZuchhipBarony struct {
	feud.BaseBarony
}

var BZuchhip叙侈波 feud.Barony = &叙侈波ZuchhipBarony{}

func init() {
    f := BZuchhip叙侈波.(*叙侈波ZuchhipBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zuchhip",
		TitleName: "叙侈波",
		TitleCode: "b_zuchhip",
	}
}
