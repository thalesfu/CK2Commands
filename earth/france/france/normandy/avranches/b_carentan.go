package avranches

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡朗唐CarentanBarony struct {
	feud.BaseBarony
}

var BCarentan卡朗唐 feud.Barony = &卡朗唐CarentanBarony{}

func init() {
    f := BCarentan卡朗唐.(*卡朗唐CarentanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carentan",
		TitleName: "卡朗唐",
		TitleCode: "b_carentan",
	}
}
