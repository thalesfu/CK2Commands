package mustang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 冲瑟尔ChhoserBarony struct {
	feud.BaseBarony
}

var BChhoser冲瑟尔 feud.Barony = &冲瑟尔ChhoserBarony{}

func init() {
    f := BChhoser冲瑟尔.(*冲瑟尔ChhoserBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chhoser",
		TitleName: "冲瑟尔",
		TitleCode: "b_chhoser",
	}
}
