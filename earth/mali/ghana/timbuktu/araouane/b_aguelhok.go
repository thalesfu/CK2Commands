package araouane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格勒霍克AguelhokBarony struct {
	feud.BaseBarony
}

var BAguelhok阿格勒霍克 feud.Barony = &阿格勒霍克AguelhokBarony{}

func init() {
	f := BAguelhok阿格勒霍克.(*阿格勒霍克AguelhokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aguelhok",
		TitleName: "阿格勒霍克",
		TitleCode: "b_aguelhok",
	}
}
