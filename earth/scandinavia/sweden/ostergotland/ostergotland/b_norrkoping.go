package ostergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 北雪平NorrkopingBarony struct {
	feud.BaseBarony
}

var BNorrkoping北雪平 feud.Barony = &北雪平NorrkopingBarony{}

func init() {
	f := BNorrkoping北雪平.(*北雪平NorrkopingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "norrkoping",
		TitleName: "北雪平",
		TitleCode: "b_norrkoping",
	}
}
