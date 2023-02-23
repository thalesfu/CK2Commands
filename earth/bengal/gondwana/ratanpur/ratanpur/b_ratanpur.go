package ratanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 剌怛那补罗RatanpurBarony struct {
	feud.BaseBarony
}

var BRatanpur剌怛那补罗 feud.Barony = &剌怛那补罗RatanpurBarony{}

func init() {
	f := BRatanpur剌怛那补罗.(*剌怛那补罗RatanpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ratanpur",
		TitleName: "剌怛那补罗",
		TitleCode: "b_ratanpur",
	}
}
