package osterreich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦格拉姆WagramBarony struct {
	feud.BaseBarony
}

var BWagram瓦格拉姆 feud.Barony = &瓦格拉姆WagramBarony{}

func init() {
	f := BWagram瓦格拉姆.(*瓦格拉姆WagramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wagram",
		TitleName: "瓦格拉姆",
		TitleCode: "b_wagram",
	}
}
