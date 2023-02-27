package faiyum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿耳西倪Arsinoe_faiyumBarony struct {
	feud.BaseBarony
}

var BArsinoe_faiyum阿耳西倪 feud.Barony = &阿耳西倪Arsinoe_faiyumBarony{}

func init() {
    f := BArsinoe_faiyum阿耳西倪.(*阿耳西倪Arsinoe_faiyumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arsinoe_faiyum",
		TitleName: "阿耳西倪",
		TitleCode: "b_arsinoe_faiyum",
	}
}
