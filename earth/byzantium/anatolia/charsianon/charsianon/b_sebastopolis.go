package charsianon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞巴斯托波利斯SebastopolisBarony struct {
	feud.BaseBarony
}

var BSebastopolis塞巴斯托波利斯 feud.Barony = &塞巴斯托波利斯SebastopolisBarony{}

func init() {
	f := BSebastopolis塞巴斯托波利斯.(*塞巴斯托波利斯SebastopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sebastopolis",
		TitleName: "塞巴斯托波利斯",
		TitleCode: "b_sebastopolis",
	}
}
