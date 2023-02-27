package balkhash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔沙特GulshatBarony struct {
	feud.BaseBarony
}

var BGulshat古尔沙特 feud.Barony = &古尔沙特GulshatBarony{}

func init() {
    f := BGulshat古尔沙特.(*古尔沙特GulshatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gulshat",
		TitleName: "古尔沙特",
		TitleCode: "b_gulshat",
	}
}
