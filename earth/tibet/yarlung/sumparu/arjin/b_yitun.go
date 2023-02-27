package arjin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 依吞YitunBarony struct {
	feud.BaseBarony
}

var BYitun依吞 feud.Barony = &依吞YitunBarony{}

func init() {
    f := BYitun依吞.(*依吞YitunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yitun",
		TitleName: "依吞",
		TitleCode: "b_yitun",
	}
}
