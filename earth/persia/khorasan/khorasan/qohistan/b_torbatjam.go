package qohistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔巴特贾姆TorbatjamBarony struct {
	feud.BaseBarony
}

var BTorbatjam托尔巴特贾姆 feud.Barony = &托尔巴特贾姆TorbatjamBarony{}

func init() {
    f := BTorbatjam托尔巴特贾姆.(*托尔巴特贾姆TorbatjamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torbatjam",
		TitleName: "托尔巴特贾姆",
		TitleCode: "b_torbatjam",
	}
}
