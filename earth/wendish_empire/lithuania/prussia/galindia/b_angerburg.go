package galindia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安格堡AngerburgBarony struct {
	feud.BaseBarony
}

var BAngerburg安格堡 feud.Barony = &安格堡AngerburgBarony{}

func init() {
    f := BAngerburg安格堡.(*安格堡AngerburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "angerburg",
		TitleName: "安格堡",
		TitleCode: "b_angerburg",
	}
}
