package alexandria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈马姆HammamBarony struct {
	feud.BaseBarony
}

var BHammam哈马姆 feud.Barony = &哈马姆HammamBarony{}

func init() {
    f := BHammam哈马姆.(*哈马姆HammamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hammam",
		TitleName: "哈马姆",
		TitleCode: "b_hammam",
	}
}
