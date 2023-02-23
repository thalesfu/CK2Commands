package hedmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯唐厄StangeBarony struct {
	feud.BaseBarony
}

var BStange斯唐厄 feud.Barony = &斯唐厄StangeBarony{}

func init() {
	f := BStange斯唐厄.(*斯唐厄StangeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stange",
		TitleName: "斯唐厄",
		TitleCode: "b_stange",
	}
}
