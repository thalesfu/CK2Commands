package tsambagarav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌雷格UuregBarony struct {
	feud.BaseBarony
}

var BUureg乌雷格 feud.Barony = &乌雷格UuregBarony{}

func init() {
	f := BUureg乌雷格.(*乌雷格UuregBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uureg",
		TitleName: "乌雷格",
		TitleCode: "b_uureg",
	}
}
