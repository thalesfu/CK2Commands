package halogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟姆纳SomnaBarony struct {
	feud.BaseBarony
}

var BSomna瑟姆纳 feud.Barony = &瑟姆纳SomnaBarony{}

func init() {
	f := BSomna瑟姆纳.(*瑟姆纳SomnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "somna",
		TitleName: "瑟姆纳",
		TitleCode: "b_somna",
	}
}
