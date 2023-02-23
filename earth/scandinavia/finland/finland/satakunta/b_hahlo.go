package satakunta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈赫洛HahloBarony struct {
	feud.BaseBarony
}

var BHahlo哈赫洛 feud.Barony = &哈赫洛HahloBarony{}

func init() {
	f := BHahlo哈赫洛.(*哈赫洛HahloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hahlo",
		TitleName: "哈赫洛",
		TitleCode: "b_hahlo",
	}
}
