package narim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔塔KoltaBarony struct {
	feud.BaseBarony
}

var BKolta科尔塔 feud.Barony = &科尔塔KoltaBarony{}

func init() {
    f := BKolta科尔塔.(*科尔塔KoltaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolta",
		TitleName: "科尔塔",
		TitleCode: "b_kolta",
	}
}
