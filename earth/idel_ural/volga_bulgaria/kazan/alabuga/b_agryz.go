package alabuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格雷兹AgryzBarony struct {
	feud.BaseBarony
}

var BAgryz阿格雷兹 feud.Barony = &阿格雷兹AgryzBarony{}

func init() {
    f := BAgryz阿格雷兹.(*阿格雷兹AgryzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agryz",
		TitleName: "阿格雷兹",
		TitleCode: "b_agryz",
	}
}
