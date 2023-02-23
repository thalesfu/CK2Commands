package radha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 揭婆KatwaBarony struct {
	feud.BaseBarony
}

var BKatwa揭婆 feud.Barony = &揭婆KatwaBarony{}

func init() {
	f := BKatwa揭婆.(*揭婆KatwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katwa",
		TitleName: "揭婆",
		TitleCode: "b_katwa",
	}
}
