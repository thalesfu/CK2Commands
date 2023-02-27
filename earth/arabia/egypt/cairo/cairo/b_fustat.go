package cairo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福斯塔特FustatBarony struct {
	feud.BaseBarony
}

var BFustat福斯塔特 feud.Barony = &福斯塔特FustatBarony{}

func init() {
    f := BFustat福斯塔特.(*福斯塔特FustatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fustat",
		TitleName: "福斯塔特",
		TitleCode: "b_fustat",
	}
}
