package markam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 左贡ZogongBarony struct {
	feud.BaseBarony
}

var BZogong左贡 feud.Barony = &左贡ZogongBarony{}

func init() {
	f := BZogong左贡.(*左贡ZogongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zogong",
		TitleName: "左贡",
		TitleCode: "b_zogong",
	}
}
