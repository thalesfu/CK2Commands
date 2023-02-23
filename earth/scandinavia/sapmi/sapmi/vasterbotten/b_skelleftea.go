package vasterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢莱夫特奥SkellefteaBarony struct {
	feud.BaseBarony
}

var BSkelleftea谢莱夫特奥 feud.Barony = &谢莱夫特奥SkellefteaBarony{}

func init() {
	f := BSkelleftea谢莱夫特奥.(*谢莱夫特奥SkellefteaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skelleftea",
		TitleName: "谢莱夫特奥",
		TitleCode: "b_skelleftea",
	}
}
