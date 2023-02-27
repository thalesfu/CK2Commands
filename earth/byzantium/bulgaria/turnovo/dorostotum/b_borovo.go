package dorostotum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博罗沃BorovoBarony struct {
	feud.BaseBarony
}

var BBorovo博罗沃 feud.Barony = &博罗沃BorovoBarony{}

func init() {
    f := BBorovo博罗沃.(*博罗沃BorovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borovo",
		TitleName: "博罗沃",
		TitleCode: "b_borovo",
	}
}
