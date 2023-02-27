package cyrenaica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拱北AlqubbahBarony struct {
	feud.BaseBarony
}

var BAlqubbah拱北 feud.Barony = &拱北AlqubbahBarony{}

func init() {
    f := BAlqubbah拱北.(*拱北AlqubbahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alqubbah",
		TitleName: "拱北",
		TitleCode: "b_alqubbah",
	}
}
