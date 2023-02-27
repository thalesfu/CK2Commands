package amastris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿马斯拉AmasraBarony struct {
	feud.BaseBarony
}

var BAmasra阿马斯拉 feud.Barony = &阿马斯拉AmasraBarony{}

func init() {
    f := BAmasra阿马斯拉.(*阿马斯拉AmasraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amasra",
		TitleName: "阿马斯拉",
		TitleCode: "b_amasra",
	}
}
