package evreux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿让唐ArgentanBarony struct {
	feud.BaseBarony
}

var BArgentan阿让唐 feud.Barony = &阿让唐ArgentanBarony{}

func init() {
	f := BArgentan阿让唐.(*阿让唐ArgentanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "argentan",
		TitleName: "阿让唐",
		TitleCode: "b_argentan",
	}
}
