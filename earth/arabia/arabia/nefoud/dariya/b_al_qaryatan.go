package dariya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖尔亚廷Al_qaryatanBarony struct {
	feud.BaseBarony
}

var BAl_qaryatan盖尔亚廷 feud.Barony = &盖尔亚廷Al_qaryatanBarony{}

func init() {
    f := BAl_qaryatan盖尔亚廷.(*盖尔亚廷Al_qaryatanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_qaryatan",
		TitleName: "盖尔亚廷",
		TitleCode: "b_al_qaryatan",
	}
}
