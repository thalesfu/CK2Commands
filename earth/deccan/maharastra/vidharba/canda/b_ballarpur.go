package canda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋剌罗补罗BallarpurBarony struct {
	feud.BaseBarony
}

var BBallarpur跋剌罗补罗 feud.Barony = &跋剌罗补罗BallarpurBarony{}

func init() {
	f := BBallarpur跋剌罗补罗.(*跋剌罗补罗BallarpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ballarpur",
		TitleName: "跋剌罗补罗",
		TitleCode: "b_ballarpur",
	}
}
