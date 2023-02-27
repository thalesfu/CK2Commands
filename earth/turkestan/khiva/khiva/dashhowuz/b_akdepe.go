package dashhowuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克杰佩AkdepeBarony struct {
	feud.BaseBarony
}

var BAkdepe阿克杰佩 feud.Barony = &阿克杰佩AkdepeBarony{}

func init() {
    f := BAkdepe阿克杰佩.(*阿克杰佩AkdepeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akdepe",
		TitleName: "阿克杰佩",
		TitleCode: "b_akdepe",
	}
}
