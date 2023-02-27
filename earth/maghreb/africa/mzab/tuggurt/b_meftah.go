package tuggurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅夫塔MeftahBarony struct {
	feud.BaseBarony
}

var BMeftah梅夫塔 feud.Barony = &梅夫塔MeftahBarony{}

func init() {
    f := BMeftah梅夫塔.(*梅夫塔MeftahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meftah",
		TitleName: "梅夫塔",
		TitleCode: "b_meftah",
	}
}
