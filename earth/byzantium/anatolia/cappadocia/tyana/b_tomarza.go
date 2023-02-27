package tyana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托马尔扎TomarzaBarony struct {
	feud.BaseBarony
}

var BTomarza托马尔扎 feud.Barony = &托马尔扎TomarzaBarony{}

func init() {
    f := BTomarza托马尔扎.(*托马尔扎TomarzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tomarza",
		TitleName: "托马尔扎",
		TitleCode: "b_tomarza",
	}
}
