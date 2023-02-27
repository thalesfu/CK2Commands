package amous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格雷GrayBarony struct {
	feud.BaseBarony
}

var BGray格雷 feud.Barony = &格雷GrayBarony{}

func init() {
    f := BGray格雷.(*格雷GrayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gray",
		TitleName: "格雷",
		TitleCode: "b_gray",
	}
}
