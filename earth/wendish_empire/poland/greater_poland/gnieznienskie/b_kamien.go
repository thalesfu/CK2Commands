package gnieznienskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡缅KamienBarony struct {
	feud.BaseBarony
}

var BKamien卡缅 feud.Barony = &卡缅KamienBarony{}

func init() {
    f := BKamien卡缅.(*卡缅KamienBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamien",
		TitleName: "卡缅",
		TitleCode: "b_kamien",
	}
}
