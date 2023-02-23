package suf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比当BidaneBarony struct {
	feud.BaseBarony
}

var BBidane比当 feud.Barony = &比当BidaneBarony{}

func init() {
	f := BBidane比当.(*比当BidaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bidane",
		TitleName: "比当",
		TitleCode: "b_bidane",
	}
}
