package medapata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 翳迦林伽视EklingjiBarony struct {
	feud.BaseBarony
}

var BEklingji翳迦林伽视 feud.Barony = &翳迦林伽视EklingjiBarony{}

func init() {
    f := BEklingji翳迦林伽视.(*翳迦林伽视EklingjiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eklingji",
		TitleName: "翳迦林伽视",
		TitleCode: "b_eklingji",
	}
}
