package saris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉拉尔特GiralthBarony struct {
	feud.BaseBarony
}

var BGiralth吉拉尔特 feud.Barony = &吉拉尔特GiralthBarony{}

func init() {
	f := BGiralth吉拉尔特.(*吉拉尔特GiralthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "giralth",
		TitleName: "吉拉尔特",
		TitleCode: "b_giralth",
	}
}
