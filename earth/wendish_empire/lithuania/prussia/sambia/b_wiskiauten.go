package sambia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维斯基奥滕WiskiautenBarony struct {
	feud.BaseBarony
}

var BWiskiauten维斯基奥滕 feud.Barony = &维斯基奥滕WiskiautenBarony{}

func init() {
    f := BWiskiauten维斯基奥滕.(*维斯基奥滕WiskiautenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wiskiauten",
		TitleName: "维斯基奥滕",
		TitleCode: "b_wiskiauten",
	}
}
