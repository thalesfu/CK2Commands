package sarysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茹尔德兹ZhuldyzBarony struct {
	feud.BaseBarony
}

var BZhuldyz茹尔德兹 feud.Barony = &茹尔德兹ZhuldyzBarony{}

func init() {
	f := BZhuldyz茹尔德兹.(*茹尔德兹ZhuldyzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhuldyz",
		TitleName: "茹尔德兹",
		TitleCode: "b_zhuldyz",
	}
}
