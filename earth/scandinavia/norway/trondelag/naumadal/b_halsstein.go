package naumadal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔斯泰因HalssteinBarony struct {
	feud.BaseBarony
}

var BHalsstein哈尔斯泰因 feud.Barony = &哈尔斯泰因HalssteinBarony{}

func init() {
	f := BHalsstein哈尔斯泰因.(*哈尔斯泰因HalssteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halsstein",
		TitleName: "哈尔斯泰因",
		TitleCode: "b_halsstein",
	}
}
