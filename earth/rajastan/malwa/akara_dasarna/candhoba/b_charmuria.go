package candhoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 遮罗牟梨耶CharmuriaBarony struct {
	feud.BaseBarony
}

var BCharmuria遮罗牟梨耶 feud.Barony = &遮罗牟梨耶CharmuriaBarony{}

func init() {
    f := BCharmuria遮罗牟梨耶.(*遮罗牟梨耶CharmuriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charmuria",
		TitleName: "遮罗牟梨耶",
		TitleCode: "b_charmuria",
	}
}
