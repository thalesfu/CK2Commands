package birlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯尔拉德BarladBarony struct {
	feud.BaseBarony
}

var BBarlad伯尔拉德 feud.Barony = &伯尔拉德BarladBarony{}

func init() {
    f := BBarlad伯尔拉德.(*伯尔拉德BarladBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barlad",
		TitleName: "伯尔拉德",
		TitleCode: "b_barlad",
	}
}
