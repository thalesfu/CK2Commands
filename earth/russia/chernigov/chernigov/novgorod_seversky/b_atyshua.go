package novgorod_seversky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿秋沙AtyshuaBarony struct {
	feud.BaseBarony
}

var BAtyshua阿秋沙 feud.Barony = &阿秋沙AtyshuaBarony{}

func init() {
    f := BAtyshua阿秋沙.(*阿秋沙AtyshuaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atyshua",
		TitleName: "阿秋沙",
		TitleCode: "b_atyshua",
	}
}
