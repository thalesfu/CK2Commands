package maldives

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孚瓦姆拉FuvahmulahBarony struct {
	feud.BaseBarony
}

var BFuvahmulah孚瓦姆拉 feud.Barony = &孚瓦姆拉FuvahmulahBarony{}

func init() {
    f := BFuvahmulah孚瓦姆拉.(*孚瓦姆拉FuvahmulahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fuvahmulah",
		TitleName: "孚瓦姆拉",
		TitleCode: "b_fuvahmulah",
	}
}
