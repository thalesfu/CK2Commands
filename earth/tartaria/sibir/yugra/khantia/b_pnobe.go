package khantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普诺比耶PnobeBarony struct {
	feud.BaseBarony
}

var BPnobe普诺比耶 feud.Barony = &普诺比耶PnobeBarony{}

func init() {
    f := BPnobe普诺比耶.(*普诺比耶PnobeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pnobe",
		TitleName: "普诺比耶",
		TitleCode: "b_pnobe",
	}
}
