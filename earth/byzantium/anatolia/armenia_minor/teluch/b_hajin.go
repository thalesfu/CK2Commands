package teluch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈金HajinBarony struct {
	feud.BaseBarony
}

var BHajin哈金 feud.Barony = &哈金HajinBarony{}

func init() {
    f := BHajin哈金.(*哈金HajinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hajin",
		TitleName: "哈金",
		TitleCode: "b_hajin",
	}
}
