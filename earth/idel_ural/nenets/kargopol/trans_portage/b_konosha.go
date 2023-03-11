package trans_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科诺沙KonoshaBarony struct {
	feud.BaseBarony
}

var BKonosha科诺沙 feud.Barony = &科诺沙KonoshaBarony{}

func init() {
    f := BKonosha科诺沙.(*科诺沙KonoshaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konosha",
		TitleName: "科诺沙",
		TitleCode: "b_konosha",
	}
}
