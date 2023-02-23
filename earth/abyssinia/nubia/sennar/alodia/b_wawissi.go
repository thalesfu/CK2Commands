package alodia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦维斯WawissiBarony struct {
	feud.BaseBarony
}

var BWawissi瓦维斯 feud.Barony = &瓦维斯WawissiBarony{}

func init() {
	f := BWawissi瓦维斯.(*瓦维斯WawissiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wawissi",
		TitleName: "瓦维斯",
		TitleCode: "b_wawissi",
	}
}
