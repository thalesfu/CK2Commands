package soli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布罗德BrodBarony struct {
	feud.BaseBarony
}

var BBrod布罗德 feud.Barony = &布罗德BrodBarony{}

func init() {
	f := BBrod布罗德.(*布罗德BrodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brod",
		TitleName: "布罗德",
		TitleCode: "b_brod",
	}
}
