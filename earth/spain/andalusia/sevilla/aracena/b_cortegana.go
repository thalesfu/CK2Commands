package aracena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔特加纳CorteganaBarony struct {
	feud.BaseBarony
}

var BCortegana科尔特加纳 feud.Barony = &科尔特加纳CorteganaBarony{}

func init() {
	f := BCortegana科尔特加纳.(*科尔特加纳CorteganaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cortegana",
		TitleName: "科尔特加纳",
		TitleCode: "b_cortegana",
	}
}
