package genoa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基亚瓦里ChiavariBarony struct {
	feud.BaseBarony
}

var BChiavari基亚瓦里 feud.Barony = &基亚瓦里ChiavariBarony{}

func init() {
    f := BChiavari基亚瓦里.(*基亚瓦里ChiavariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chiavari",
		TitleName: "基亚瓦里",
		TitleCode: "b_chiavari",
	}
}
