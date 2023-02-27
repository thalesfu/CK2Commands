package hellas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安菲萨AmphissaBarony struct {
	feud.BaseBarony
}

var BAmphissa安菲萨 feud.Barony = &安菲萨AmphissaBarony{}

func init() {
    f := BAmphissa安菲萨.(*安菲萨AmphissaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amphissa",
		TitleName: "安菲萨",
		TitleCode: "b_amphissa",
	}
}
