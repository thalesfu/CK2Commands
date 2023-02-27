package tyrus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔福SyrbelfortBarony struct {
	feud.BaseBarony
}

var BSyrbelfort贝尔福 feud.Barony = &贝尔福SyrbelfortBarony{}

func init() {
    f := BSyrbelfort贝尔福.(*贝尔福SyrbelfortBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syrbelfort",
		TitleName: "贝尔福",
		TitleCode: "b_syrbelfort",
	}
}
