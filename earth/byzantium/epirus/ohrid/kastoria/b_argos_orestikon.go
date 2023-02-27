package kastoria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔戈斯奥雷斯蒂孔Argos_orestikonBarony struct {
	feud.BaseBarony
}

var BArgos_orestikon阿尔戈斯奥雷斯蒂孔 feud.Barony = &阿尔戈斯奥雷斯蒂孔Argos_orestikonBarony{}

func init() {
    f := BArgos_orestikon阿尔戈斯奥雷斯蒂孔.(*阿尔戈斯奥雷斯蒂孔Argos_orestikonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "argos_orestikon",
		TitleName: "阿尔戈斯奥雷斯蒂孔",
		TitleCode: "b_argos_orestikon",
	}
}
