package tigrinya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆巴马德里亚Amba_maderiyaBarony struct {
	feud.BaseBarony
}

var BAmba_maderiya阿姆巴马德里亚 feud.Barony = &阿姆巴马德里亚Amba_maderiyaBarony{}

func init() {
    f := BAmba_maderiya阿姆巴马德里亚.(*阿姆巴马德里亚Amba_maderiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amba_maderiya",
		TitleName: "阿姆巴马德里亚",
		TitleCode: "b_amba_maderiya",
	}
}
