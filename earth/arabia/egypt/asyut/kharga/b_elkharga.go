package kharga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈里杰ElkhargaBarony struct {
	feud.BaseBarony
}

var BElkharga哈里杰 feud.Barony = &哈里杰ElkhargaBarony{}

func init() {
    f := BElkharga哈里杰.(*哈里杰ElkhargaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elkharga",
		TitleName: "哈里杰",
		TitleCode: "b_elkharga",
	}
}
