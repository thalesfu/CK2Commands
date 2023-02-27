package tannu_ola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 唐麓岭Tannu_olaBarony struct {
	feud.BaseBarony
}

var BTannu_ola唐麓岭 feud.Barony = &唐麓岭Tannu_olaBarony{}

func init() {
    f := BTannu_ola唐麓岭.(*唐麓岭Tannu_olaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tannu_ola",
		TitleName: "唐麓岭",
		TitleCode: "b_tannu_ola",
	}
}
