package leventina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝林佐纳BellinzonaBarony struct {
	feud.BaseBarony
}

var BBellinzona贝林佐纳 feud.Barony = &贝林佐纳BellinzonaBarony{}

func init() {
    f := BBellinzona贝林佐纳.(*贝林佐纳BellinzonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bellinzona",
		TitleName: "贝林佐纳",
		TitleCode: "b_bellinzona",
	}
}
