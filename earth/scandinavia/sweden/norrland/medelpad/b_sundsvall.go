package medelpad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松兹瓦尔SundsvallBarony struct {
	feud.BaseBarony
}

var BSundsvall松兹瓦尔 feud.Barony = &松兹瓦尔SundsvallBarony{}

func init() {
    f := BSundsvall松兹瓦尔.(*松兹瓦尔SundsvallBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sundsvall",
		TitleName: "松兹瓦尔",
		TitleCode: "b_sundsvall",
	}
}
