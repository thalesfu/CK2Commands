package ikh_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 也客孛黑答Ikh_bogdBarony struct {
	feud.BaseBarony
}

var BIkh_bogd也客孛黑答 feud.Barony = &也客孛黑答Ikh_bogdBarony{}

func init() {
    f := BIkh_bogd也客孛黑答.(*也客孛黑答Ikh_bogdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ikh_bogd",
		TitleName: "也客孛黑答",
		TitleCode: "b_ikh_bogd",
	}
}
