package powys

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉斯库姆GlascwmBarony struct {
	feud.BaseBarony
}

var BGlascwm格拉斯库姆 feud.Barony = &格拉斯库姆GlascwmBarony{}

func init() {
	f := BGlascwm格拉斯库姆.(*格拉斯库姆GlascwmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glascwm",
		TitleName: "格拉斯库姆",
		TitleCode: "b_glascwm",
	}
}
