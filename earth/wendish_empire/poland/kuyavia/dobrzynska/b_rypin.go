package dobrzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷平RypinBarony struct {
	feud.BaseBarony
}

var BRypin雷平 feud.Barony = &雷平RypinBarony{}

func init() {
    f := BRypin雷平.(*雷平RypinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rypin",
		TitleName: "雷平",
		TitleCode: "b_rypin",
	}
}
