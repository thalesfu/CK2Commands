package trondelag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯韦勒斯堡SverresborgBarony struct {
	feud.BaseBarony
}

var BSverresborg斯韦勒斯堡 feud.Barony = &斯韦勒斯堡SverresborgBarony{}

func init() {
	f := BSverresborg斯韦勒斯堡.(*斯韦勒斯堡SverresborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sverresborg",
		TitleName: "斯韦勒斯堡",
		TitleCode: "b_sverresborg",
	}
}
