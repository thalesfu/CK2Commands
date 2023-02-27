package khotan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 精绝JingjueBarony struct {
	feud.BaseBarony
}

var BJingjue精绝 feud.Barony = &精绝JingjueBarony{}

func init() {
    f := BJingjue精绝.(*精绝JingjueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jingjue",
		TitleName: "精绝",
		TitleCode: "b_jingjue",
	}
}
