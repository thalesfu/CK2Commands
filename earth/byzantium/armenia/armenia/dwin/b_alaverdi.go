package dwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉韦尔迪AlaverdiBarony struct {
	feud.BaseBarony
}

var BAlaverdi阿拉韦尔迪 feud.Barony = &阿拉韦尔迪AlaverdiBarony{}

func init() {
    f := BAlaverdi阿拉韦尔迪.(*阿拉韦尔迪AlaverdiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alaverdi",
		TitleName: "阿拉韦尔迪",
		TitleCode: "b_alaverdi",
	}
}
