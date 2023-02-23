package trencin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊洛沃IllavaBarony struct {
	feud.BaseBarony
}

var BIllava伊洛沃 feud.Barony = &伊洛沃IllavaBarony{}

func init() {
	f := BIllava伊洛沃.(*伊洛沃IllavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "illava",
		TitleName: "伊洛沃",
		TitleCode: "b_illava",
	}
}
