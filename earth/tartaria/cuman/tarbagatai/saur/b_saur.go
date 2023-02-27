package saur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨吾尔SaurBarony struct {
	feud.BaseBarony
}

var BSaur萨吾尔 feud.Barony = &萨吾尔SaurBarony{}

func init() {
    f := BSaur萨吾尔.(*萨吾尔SaurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saur",
		TitleName: "萨吾尔",
		TitleCode: "b_saur",
	}
}
