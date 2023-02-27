package gezira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 唐加西TangasiBarony struct {
	feud.BaseBarony
}

var BTangasi唐加西 feud.Barony = &唐加西TangasiBarony{}

func init() {
    f := BTangasi唐加西.(*唐加西TangasiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tangasi",
		TitleName: "唐加西",
		TitleCode: "b_tangasi",
	}
}
