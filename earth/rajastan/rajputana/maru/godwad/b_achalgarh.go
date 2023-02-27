package godwad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿遮罗姞利呬AchalgarhBarony struct {
	feud.BaseBarony
}

var BAchalgarh阿遮罗姞利呬 feud.Barony = &阿遮罗姞利呬AchalgarhBarony{}

func init() {
    f := BAchalgarh阿遮罗姞利呬.(*阿遮罗姞利呬AchalgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "achalgarh",
		TitleName: "阿遮罗姞利呬",
		TitleCode: "b_achalgarh",
	}
}
