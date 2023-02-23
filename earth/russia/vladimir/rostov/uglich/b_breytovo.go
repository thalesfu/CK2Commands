package uglich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布列伊托沃BreytovoBarony struct {
	feud.BaseBarony
}

var BBreytovo布列伊托沃 feud.Barony = &布列伊托沃BreytovoBarony{}

func init() {
	f := BBreytovo布列伊托沃.(*布列伊托沃BreytovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "breytovo",
		TitleName: "布列伊托沃",
		TitleCode: "b_breytovo",
	}
}
