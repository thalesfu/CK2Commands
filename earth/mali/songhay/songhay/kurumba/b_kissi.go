package kurumba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基西KissiBarony struct {
	feud.BaseBarony
}

var BKissi基西 feud.Barony = &基西KissiBarony{}

func init() {
	f := BKissi基西.(*基西KissiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kissi",
		TitleName: "基西",
		TitleCode: "b_kissi",
	}
}
