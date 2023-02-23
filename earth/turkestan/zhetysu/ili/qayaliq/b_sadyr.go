package qayaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨德尔SadyrBarony struct {
	feud.BaseBarony
}

var BSadyr萨德尔 feud.Barony = &萨德尔SadyrBarony{}

func init() {
	f := BSadyr萨德尔.(*萨德尔SadyrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sadyr",
		TitleName: "萨德尔",
		TitleCode: "b_sadyr",
	}
}
