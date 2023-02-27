package ile_de_france

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃唐普EtampesBarony struct {
	feud.BaseBarony
}

var BEtampes埃唐普 feud.Barony = &埃唐普EtampesBarony{}

func init() {
    f := BEtampes埃唐普.(*埃唐普EtampesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "etampes",
		TitleName: "埃唐普",
		TitleCode: "b_etampes",
	}
}
