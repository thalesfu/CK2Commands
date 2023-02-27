package leeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 唐克斯特DoncasterBarony struct {
	feud.BaseBarony
}

var BDoncaster唐克斯特 feud.Barony = &唐克斯特DoncasterBarony{}

func init() {
    f := BDoncaster唐克斯特.(*唐克斯特DoncasterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doncaster",
		TitleName: "唐克斯特",
		TitleCode: "b_doncaster",
	}
}
