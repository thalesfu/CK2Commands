package gaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩提伽耶Bodh_gayaBarony struct {
	feud.BaseBarony
}

var BBodh_gaya菩提伽耶 feud.Barony = &菩提伽耶Bodh_gayaBarony{}

func init() {
    f := BBodh_gaya菩提伽耶.(*菩提伽耶Bodh_gayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bodh_gaya",
		TitleName: "菩提伽耶",
		TitleCode: "b_bodh_gaya",
	}
}
