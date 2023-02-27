package ostergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 申宁厄SkanningeBarony struct {
	feud.BaseBarony
}

var BSkanninge申宁厄 feud.Barony = &申宁厄SkanningeBarony{}

func init() {
    f := BSkanninge申宁厄.(*申宁厄SkanningeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skanninge",
		TitleName: "申宁厄",
		TitleCode: "b_skanninge",
	}
}
