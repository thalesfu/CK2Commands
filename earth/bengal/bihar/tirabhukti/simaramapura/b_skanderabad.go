package simaramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞建咥罗跋陀SkanderabadBarony struct {
	feud.BaseBarony
}

var BSkanderabad塞建咥罗跋陀 feud.Barony = &塞建咥罗跋陀SkanderabadBarony{}

func init() {
	f := BSkanderabad塞建咥罗跋陀.(*塞建咥罗跋陀SkanderabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skanderabad",
		TitleName: "塞建咥罗跋陀",
		TitleCode: "b_skanderabad",
	}
}
