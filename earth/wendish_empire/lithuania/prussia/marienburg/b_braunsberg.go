package marienburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布劳恩斯贝格BraunsbergBarony struct {
	feud.BaseBarony
}

var BBraunsberg布劳恩斯贝格 feud.Barony = &布劳恩斯贝格BraunsbergBarony{}

func init() {
    f := BBraunsberg布劳恩斯贝格.(*布劳恩斯贝格BraunsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "braunsberg",
		TitleName: "布劳恩斯贝格",
		TitleCode: "b_braunsberg",
	}
}
