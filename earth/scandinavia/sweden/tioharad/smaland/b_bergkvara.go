package smaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝里克瓦拉BergkvaraBarony struct {
	feud.BaseBarony
}

var BBergkvara贝里克瓦拉 feud.Barony = &贝里克瓦拉BergkvaraBarony{}

func init() {
	f := BBergkvara贝里克瓦拉.(*贝里克瓦拉BergkvaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bergkvara",
		TitleName: "贝里克瓦拉",
		TitleCode: "b_bergkvara",
	}
}
