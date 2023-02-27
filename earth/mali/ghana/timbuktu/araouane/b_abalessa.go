package araouane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴来萨AbalessaBarony struct {
	feud.BaseBarony
}

var BAbalessa阿巴来萨 feud.Barony = &阿巴来萨AbalessaBarony{}

func init() {
    f := BAbalessa阿巴来萨.(*阿巴来萨AbalessaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abalessa",
		TitleName: "阿巴来萨",
		TitleCode: "b_abalessa",
	}
}
