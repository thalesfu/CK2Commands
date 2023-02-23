package gloucester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔滕纳姆CheltenhamBarony struct {
	feud.BaseBarony
}

var BCheltenham切尔滕纳姆 feud.Barony = &切尔滕纳姆CheltenhamBarony{}

func init() {
	f := BCheltenham切尔滕纳姆.(*切尔滕纳姆CheltenhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cheltenham",
		TitleName: "切尔滕纳姆",
		TitleCode: "b_cheltenham",
	}
}
