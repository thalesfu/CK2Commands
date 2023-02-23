package vitebsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别申科维奇BeshankovichyBarony struct {
	feud.BaseBarony
}

var BBeshankovichy别申科维奇 feud.Barony = &别申科维奇BeshankovichyBarony{}

func init() {
	f := BBeshankovichy别申科维奇.(*别申科维奇BeshankovichyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beshankovichy",
		TitleName: "别申科维奇",
		TitleCode: "b_beshankovichy",
	}
}
