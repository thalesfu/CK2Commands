package biskra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪奥克巴SidiokbaBarony struct {
	feud.BaseBarony
}

var BSidiokba西迪奥克巴 feud.Barony = &西迪奥克巴SidiokbaBarony{}

func init() {
	f := BSidiokba西迪奥克巴.(*西迪奥克巴SidiokbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidiokba",
		TitleName: "西迪奥克巴",
		TitleCode: "b_sidiokba",
	}
}
