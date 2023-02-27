package lykandos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科曼格尼ComanageneBarony struct {
	feud.BaseBarony
}

var BComanagene科曼格尼 feud.Barony = &科曼格尼ComanageneBarony{}

func init() {
    f := BComanagene科曼格尼.(*科曼格尼ComanageneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "comanagene",
		TitleName: "科曼格尼",
		TitleCode: "b_comanagene",
	}
}
