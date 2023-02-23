package wiltshire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔兹伯里SalisburyBarony struct {
	feud.BaseBarony
}

var BSalisbury索尔兹伯里 feud.Barony = &索尔兹伯里SalisburyBarony{}

func init() {
	f := BSalisbury索尔兹伯里.(*索尔兹伯里SalisburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salisbury",
		TitleName: "索尔兹伯里",
		TitleCode: "b_salisbury",
	}
}
