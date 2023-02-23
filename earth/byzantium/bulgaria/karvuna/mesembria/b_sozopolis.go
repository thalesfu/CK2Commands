package mesembria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索佐波利斯SozopolisBarony struct {
	feud.BaseBarony
}

var BSozopolis索佐波利斯 feud.Barony = &索佐波利斯SozopolisBarony{}

func init() {
	f := BSozopolis索佐波利斯.(*索佐波利斯SozopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sozopolis",
		TitleName: "索佐波利斯",
		TitleCode: "b_sozopolis",
	}
}
