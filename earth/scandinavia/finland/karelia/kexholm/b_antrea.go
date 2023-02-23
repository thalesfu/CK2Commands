package kexholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安特雷亚AntreaBarony struct {
	feud.BaseBarony
}

var BAntrea安特雷亚 feud.Barony = &安特雷亚AntreaBarony{}

func init() {
	f := BAntrea安特雷亚.(*安特雷亚AntreaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "antrea",
		TitleName: "安特雷亚",
		TitleCode: "b_antrea",
	}
}
