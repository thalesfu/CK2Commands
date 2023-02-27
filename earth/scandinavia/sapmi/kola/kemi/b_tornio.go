package kemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔尼奥TornioBarony struct {
	feud.BaseBarony
}

var BTornio托尔尼奥 feud.Barony = &托尔尼奥TornioBarony{}

func init() {
    f := BTornio托尔尼奥.(*托尔尼奥TornioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tornio",
		TitleName: "托尔尼奥",
		TitleCode: "b_tornio",
	}
}
