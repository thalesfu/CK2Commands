package barkul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旧县镇JiuxianBarony struct {
	feud.BaseBarony
}

var BJiuxian旧县镇 feud.Barony = &旧县镇JiuxianBarony{}

func init() {
	f := BJiuxian旧县镇.(*旧县镇JiuxianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jiuxian",
		TitleName: "旧县镇",
		TitleCode: "b_jiuxian",
	}
}
