package tummana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 堵尔布恩TualbungBarony struct {
	feud.BaseBarony
}

var BTualbung堵尔布恩 feud.Barony = &堵尔布恩TualbungBarony{}

func init() {
	f := BTualbung堵尔布恩.(*堵尔布恩TualbungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tualbung",
		TitleName: "堵尔布恩",
		TitleCode: "b_tualbung",
	}
}
