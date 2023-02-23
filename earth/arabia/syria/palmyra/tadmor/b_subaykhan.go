package tadmor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏拜汉SubaykhanBarony struct {
	feud.BaseBarony
}

var BSubaykhan苏拜汉 feud.Barony = &苏拜汉SubaykhanBarony{}

func init() {
	f := BSubaykhan苏拜汉.(*苏拜汉SubaykhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "subaykhan",
		TitleName: "苏拜汉",
		TitleCode: "b_subaykhan",
	}
}
