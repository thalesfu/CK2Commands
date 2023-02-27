package homs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞代德SadadBarony struct {
	feud.BaseBarony
}

var BSadad塞代德 feud.Barony = &塞代德SadadBarony{}

func init() {
    f := BSadad塞代德.(*塞代德SadadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sadad",
		TitleName: "塞代德",
		TitleCode: "b_sadad",
	}
}
