package east_karelia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢戈泽罗SegozeroBarony struct {
	feud.BaseBarony
}

var BSegozero谢戈泽罗 feud.Barony = &谢戈泽罗SegozeroBarony{}

func init() {
    f := BSegozero谢戈泽罗.(*谢戈泽罗SegozeroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "segozero",
		TitleName: "谢戈泽罗",
		TitleCode: "b_segozero",
	}
}
