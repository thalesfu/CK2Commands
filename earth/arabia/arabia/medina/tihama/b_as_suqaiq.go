package tihama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏盖格As_suqaiqBarony struct {
	feud.BaseBarony
}

var BAs_suqaiq苏盖格 feud.Barony = &苏盖格As_suqaiqBarony{}

func init() {
    f := BAs_suqaiq苏盖格.(*苏盖格As_suqaiqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "as_suqaiq",
		TitleName: "苏盖格",
		TitleCode: "b_as_suqaiq",
	}
}
