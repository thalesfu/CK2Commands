package kangra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 屈露KulluBarony struct {
	feud.BaseBarony
}

var BKullu屈露 feud.Barony = &屈露KulluBarony{}

func init() {
    f := BKullu屈露.(*屈露KulluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kullu",
		TitleName: "屈露",
		TitleCode: "b_kullu",
	}
}
