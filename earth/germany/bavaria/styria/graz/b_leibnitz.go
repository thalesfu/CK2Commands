package graz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱布尼茨LeibnitzBarony struct {
	feud.BaseBarony
}

var BLeibnitz莱布尼茨 feud.Barony = &莱布尼茨LeibnitzBarony{}

func init() {
	f := BLeibnitz莱布尼茨.(*莱布尼茨LeibnitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leibnitz",
		TitleName: "莱布尼茨",
		TitleCode: "b_leibnitz",
	}
}
