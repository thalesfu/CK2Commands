package uchturpan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赵家嘴ZhaojiazuiBarony struct {
	feud.BaseBarony
}

var BZhaojiazui赵家嘴 feud.Barony = &赵家嘴ZhaojiazuiBarony{}

func init() {
    f := BZhaojiazui赵家嘴.(*赵家嘴ZhaojiazuiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhaojiazui",
		TitleName: "赵家嘴",
		TitleCode: "b_zhaojiazui",
	}
}
