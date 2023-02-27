package mesembria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔纳VarnaBarony struct {
	feud.BaseBarony
}

var BVarna瓦尔纳 feud.Barony = &瓦尔纳VarnaBarony{}

func init() {
    f := BVarna瓦尔纳.(*瓦尔纳VarnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varna",
		TitleName: "瓦尔纳",
		TitleCode: "b_varna",
	}
}
