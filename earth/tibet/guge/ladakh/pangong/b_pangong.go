package pangong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班公PangongBarony struct {
	feud.BaseBarony
}

var BPangong班公 feud.Barony = &班公PangongBarony{}

func init() {
    f := BPangong班公.(*班公PangongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pangong",
		TitleName: "班公",
		TitleCode: "b_pangong",
	}
}
