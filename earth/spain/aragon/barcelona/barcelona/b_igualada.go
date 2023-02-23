package barcelona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊瓜拉达IgualadaBarony struct {
	feud.BaseBarony
}

var BIgualada伊瓜拉达 feud.Barony = &伊瓜拉达IgualadaBarony{}

func init() {
	f := BIgualada伊瓜拉达.(*伊瓜拉达IgualadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "igualada",
		TitleName: "伊瓜拉达",
		TitleCode: "b_igualada",
	}
}
