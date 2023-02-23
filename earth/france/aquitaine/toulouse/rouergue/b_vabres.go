package rouergue

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦布尔VabresBarony struct {
	feud.BaseBarony
}

var BVabres瓦布尔 feud.Barony = &瓦布尔VabresBarony{}

func init() {
	f := BVabres瓦布尔.(*瓦布尔VabresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vabres",
		TitleName: "瓦布尔",
		TitleCode: "b_vabres",
	}
}
