package ifni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 唐图勒TharassetBarony struct {
	feud.BaseBarony
}

var BTharasset唐图勒 feud.Barony = &唐图勒TharassetBarony{}

func init() {
	f := BTharasset唐图勒.(*唐图勒TharassetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tharasset",
		TitleName: "唐图勒",
		TitleCode: "b_tharasset",
	}
}
