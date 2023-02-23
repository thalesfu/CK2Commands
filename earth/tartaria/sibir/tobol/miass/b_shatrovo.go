package miass

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙特罗沃ShatrovoBarony struct {
	feud.BaseBarony
}

var BShatrovo沙特罗沃 feud.Barony = &沙特罗沃ShatrovoBarony{}

func init() {
	f := BShatrovo沙特罗沃.(*沙特罗沃ShatrovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shatrovo",
		TitleName: "沙特罗沃",
		TitleCode: "b_shatrovo",
	}
}
