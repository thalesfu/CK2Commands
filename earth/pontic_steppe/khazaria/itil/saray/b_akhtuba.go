package saray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赫图巴AkhtubaBarony struct {
	feud.BaseBarony
}

var BAkhtuba阿赫图巴 feud.Barony = &阿赫图巴AkhtubaBarony{}

func init() {
    f := BAkhtuba阿赫图巴.(*阿赫图巴AkhtubaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akhtuba",
		TitleName: "阿赫图巴",
		TitleCode: "b_akhtuba",
	}
}
