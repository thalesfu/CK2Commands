package somnath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃罗卡MoherakBarony struct {
	feud.BaseBarony
}

var BMoherak摩诃罗卡 feud.Barony = &摩诃罗卡MoherakBarony{}

func init() {
    f := BMoherak摩诃罗卡.(*摩诃罗卡MoherakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moherak",
		TitleName: "摩诃罗卡",
		TitleCode: "b_moherak",
	}
}
