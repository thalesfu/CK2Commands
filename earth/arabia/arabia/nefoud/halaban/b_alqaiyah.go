package halaban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加伊耶AlqaiyahBarony struct {
	feud.BaseBarony
}

var BAlqaiyah加伊耶 feud.Barony = &加伊耶AlqaiyahBarony{}

func init() {
    f := BAlqaiyah加伊耶.(*加伊耶AlqaiyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alqaiyah",
		TitleName: "加伊耶",
		TitleCode: "b_alqaiyah",
	}
}
