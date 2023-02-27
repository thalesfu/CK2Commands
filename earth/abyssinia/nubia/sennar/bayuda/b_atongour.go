package bayuda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿通古尔AtongourBarony struct {
	feud.BaseBarony
}

var BAtongour阿通古尔 feud.Barony = &阿通古尔AtongourBarony{}

func init() {
    f := BAtongour阿通古尔.(*阿通古尔AtongourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atongour",
		TitleName: "阿通古尔",
		TitleCode: "b_atongour",
	}
}
