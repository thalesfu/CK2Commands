package chagai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝希PeshaiBarony struct {
	feud.BaseBarony
}

var BPeshai贝希 feud.Barony = &贝希PeshaiBarony{}

func init() {
	f := BPeshai贝希.(*贝希PeshaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peshai",
		TitleName: "贝希",
		TitleCode: "b_peshai",
	}
}
