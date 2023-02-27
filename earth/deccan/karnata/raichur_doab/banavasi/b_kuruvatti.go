package banavasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 俱卢伐底KuruvattiBarony struct {
	feud.BaseBarony
}

var BKuruvatti俱卢伐底 feud.Barony = &俱卢伐底KuruvattiBarony{}

func init() {
    f := BKuruvatti俱卢伐底.(*俱卢伐底KuruvattiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuruvatti",
		TitleName: "俱卢伐底",
		TitleCode: "b_kuruvatti",
	}
}
