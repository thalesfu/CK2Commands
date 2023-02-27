package kongu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄卢德ErodeBarony struct {
	feud.BaseBarony
}

var BErode厄卢德 feud.Barony = &厄卢德ErodeBarony{}

func init() {
    f := BErode厄卢德.(*厄卢德ErodeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erode",
		TitleName: "厄卢德",
		TitleCode: "b_erode",
	}
}
