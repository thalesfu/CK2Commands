package chuvash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐维利斯克TsivilskBarony struct {
	feud.BaseBarony
}

var BTsivilsk齐维利斯克 feud.Barony = &齐维利斯克TsivilskBarony{}

func init() {
    f := BTsivilsk齐维利斯克.(*齐维利斯克TsivilskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsivilsk",
		TitleName: "齐维利斯克",
		TitleCode: "b_tsivilsk",
	}
}
