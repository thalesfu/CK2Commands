package altmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 唐格明德TangermundeBarony struct {
	feud.BaseBarony
}

var BTangermunde唐格明德 feud.Barony = &唐格明德TangermundeBarony{}

func init() {
	f := BTangermunde唐格明德.(*唐格明德TangermundeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tangermunde",
		TitleName: "唐格明德",
		TitleCode: "b_tangermunde",
	}
}
