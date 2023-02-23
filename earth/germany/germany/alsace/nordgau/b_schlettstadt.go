package nordgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施莱特施塔特SchlettstadtBarony struct {
	feud.BaseBarony
}

var BSchlettstadt施莱特施塔特 feud.Barony = &施莱特施塔特SchlettstadtBarony{}

func init() {
	f := BSchlettstadt施莱特施塔特.(*施莱特施塔特SchlettstadtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schlettstadt",
		TitleName: "施莱特施塔特",
		TitleCode: "b_schlettstadt",
	}
}
