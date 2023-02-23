package turkestan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克斯佩AkespeBarony struct {
	feud.BaseBarony
}

var BAkespe阿克斯佩 feud.Barony = &阿克斯佩AkespeBarony{}

func init() {
	f := BAkespe阿克斯佩.(*阿克斯佩AkespeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akespe",
		TitleName: "阿克斯佩",
		TitleCode: "b_akespe",
	}
}
