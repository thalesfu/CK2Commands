package chuvash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉特里AlatyrBarony struct {
	feud.BaseBarony
}

var BAlatyr阿拉特里 feud.Barony = &阿拉特里AlatyrBarony{}

func init() {
    f := BAlatyr阿拉特里.(*阿拉特里AlatyrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alatyr",
		TitleName: "阿拉特里",
		TitleCode: "b_alatyr",
	}
}
