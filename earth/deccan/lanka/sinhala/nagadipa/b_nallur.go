package nagadipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 讷卢尔NallurBarony struct {
	feud.BaseBarony
}

var BNallur讷卢尔 feud.Barony = &讷卢尔NallurBarony{}

func init() {
	f := BNallur讷卢尔.(*讷卢尔NallurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nallur",
		TitleName: "讷卢尔",
		TitleCode: "b_nallur",
	}
}
