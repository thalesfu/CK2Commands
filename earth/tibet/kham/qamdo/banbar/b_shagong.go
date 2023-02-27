package banbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙贡ShagongBarony struct {
	feud.BaseBarony
}

var BShagong沙贡 feud.Barony = &沙贡ShagongBarony{}

func init() {
    f := BShagong沙贡.(*沙贡ShagongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shagong",
		TitleName: "沙贡",
		TitleCode: "b_shagong",
	}
}
