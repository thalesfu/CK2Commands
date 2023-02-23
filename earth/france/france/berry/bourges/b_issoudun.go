package bourges

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊苏丹IssoudunBarony struct {
	feud.BaseBarony
}

var BIssoudun伊苏丹 feud.Barony = &伊苏丹IssoudunBarony{}

func init() {
	f := BIssoudun伊苏丹.(*伊苏丹IssoudunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "issoudun",
		TitleName: "伊苏丹",
		TitleCode: "b_issoudun",
	}
}
