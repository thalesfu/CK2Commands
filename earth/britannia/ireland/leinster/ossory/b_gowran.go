package ossory

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 高兰GowranBarony struct {
	feud.BaseBarony
}

var BGowran高兰 feud.Barony = &高兰GowranBarony{}

func init() {
    f := BGowran高兰.(*高兰GowranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gowran",
		TitleName: "高兰",
		TitleCode: "b_gowran",
	}
}
