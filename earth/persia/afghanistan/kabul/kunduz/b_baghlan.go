package kunduz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 缚伽浪BaghlanBarony struct {
	feud.BaseBarony
}

var BBaghlan缚伽浪 feud.Barony = &缚伽浪BaghlanBarony{}

func init() {
	f := BBaghlan缚伽浪.(*缚伽浪BaghlanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baghlan",
		TitleName: "缚伽浪",
		TitleCode: "b_baghlan",
	}
}
