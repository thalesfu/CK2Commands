package hainaut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦朗谢讷ValenciennesBarony struct {
	feud.BaseBarony
}

var BValenciennes瓦朗谢讷 feud.Barony = &瓦朗谢讷ValenciennesBarony{}

func init() {
	f := BValenciennes瓦朗谢讷.(*瓦朗谢讷ValenciennesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valenciennes",
		TitleName: "瓦朗谢讷",
		TitleCode: "b_valenciennes",
	}
}
