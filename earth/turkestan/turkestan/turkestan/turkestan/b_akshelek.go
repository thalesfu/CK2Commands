package turkestan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克舍列克AkshelekBarony struct {
	feud.BaseBarony
}

var BAkshelek阿克舍列克 feud.Barony = &阿克舍列克AkshelekBarony{}

func init() {
	f := BAkshelek阿克舍列克.(*阿克舍列克AkshelekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akshelek",
		TitleName: "阿克舍列克",
		TitleCode: "b_akshelek",
	}
}
