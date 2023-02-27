package ashli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚曼套YamantawBarony struct {
	feud.BaseBarony
}

var BYamantaw亚曼套 feud.Barony = &亚曼套YamantawBarony{}

func init() {
    f := BYamantaw亚曼套.(*亚曼套YamantawBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yamantaw",
		TitleName: "亚曼套",
		TitleCode: "b_yamantaw",
	}
}
