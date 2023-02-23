package iasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅西IasiBarony struct {
	feud.BaseBarony
}

var BIasi雅西 feud.Barony = &雅西IasiBarony{}

func init() {
	f := BIasi雅西.(*雅西IasiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iasi",
		TitleName: "雅西",
		TitleCode: "b_iasi",
	}
}
