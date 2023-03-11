package kursk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普里亚米齐诺PryamitsynoBarony struct {
	feud.BaseBarony
}

var BPryamitsyno普里亚米齐诺 feud.Barony = &普里亚米齐诺PryamitsynoBarony{}

func init() {
    f := BPryamitsyno普里亚米齐诺.(*普里亚米齐诺PryamitsynoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pryamitsyno",
		TitleName: "普里亚米齐诺",
		TitleCode: "b_pryamitsyno",
	}
}
