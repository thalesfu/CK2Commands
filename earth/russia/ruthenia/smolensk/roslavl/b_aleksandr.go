package roslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚历山大AleksandrBarony struct {
	feud.BaseBarony
}

var BAleksandr亚历山大 feud.Barony = &亚历山大AleksandrBarony{}

func init() {
    f := BAleksandr亚历山大.(*亚历山大AleksandrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aleksandr",
		TitleName: "亚历山大",
		TitleCode: "b_aleksandr",
	}
}
