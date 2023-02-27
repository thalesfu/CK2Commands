package utva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克套AktauBarony struct {
	feud.BaseBarony
}

var BAktau阿克套 feud.Barony = &阿克套AktauBarony{}

func init() {
    f := BAktau阿克套.(*阿克套AktauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aktau",
		TitleName: "阿克套",
		TitleCode: "b_aktau",
	}
}
