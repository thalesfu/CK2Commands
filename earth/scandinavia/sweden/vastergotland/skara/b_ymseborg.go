package skara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于姆瑟堡YmseborgBarony struct {
	feud.BaseBarony
}

var BYmseborg于姆瑟堡 feud.Barony = &于姆瑟堡YmseborgBarony{}

func init() {
	f := BYmseborg于姆瑟堡.(*于姆瑟堡YmseborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ymseborg",
		TitleName: "于姆瑟堡",
		TitleCode: "b_ymseborg",
	}
}
