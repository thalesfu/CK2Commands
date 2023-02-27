package al_amarah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈里HaiBarony struct {
	feud.BaseBarony
}

var BHai哈里 feud.Barony = &哈里HaiBarony{}

func init() {
    f := BHai哈里.(*哈里HaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hai",
		TitleName: "哈里",
		TitleCode: "b_hai",
	}
}
