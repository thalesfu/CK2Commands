package sozopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索佐波利斯SouzopolisBarony struct {
	feud.BaseBarony
}

var BSouzopolis索佐波利斯 feud.Barony = &索佐波利斯SouzopolisBarony{}

func init() {
	f := BSouzopolis索佐波利斯.(*索佐波利斯SouzopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "souzopolis",
		TitleName: "索佐波利斯",
		TitleCode: "b_souzopolis",
	}
}
