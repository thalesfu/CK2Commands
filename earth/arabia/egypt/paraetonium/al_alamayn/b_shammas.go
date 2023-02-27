package al_alamayn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙姆ShammasBarony struct {
	feud.BaseBarony
}

var BShammas沙姆 feud.Barony = &沙姆ShammasBarony{}

func init() {
    f := BShammas沙姆.(*沙姆ShammasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shammas",
		TitleName: "沙姆",
		TitleCode: "b_shammas",
	}
}
