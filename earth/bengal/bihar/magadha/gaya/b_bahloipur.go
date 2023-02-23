package gaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆诃卢罗补罗BahloipurBarony struct {
	feud.BaseBarony
}

var BBahloipur婆诃卢罗补罗 feud.Barony = &婆诃卢罗补罗BahloipurBarony{}

func init() {
	f := BBahloipur婆诃卢罗补罗.(*婆诃卢罗补罗BahloipurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bahloipur",
		TitleName: "婆诃卢罗补罗",
		TitleCode: "b_bahloipur",
	}
}
