package amida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃塔尤拉特拉霍IdtodyolatalohoBarony struct {
	feud.BaseBarony
}

var BIdtodyolataloho埃塔尤拉特拉霍 feud.Barony = &埃塔尤拉特拉霍IdtodyolatalohoBarony{}

func init() {
	f := BIdtodyolataloho埃塔尤拉特拉霍.(*埃塔尤拉特拉霍IdtodyolatalohoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "idtodyolataloho",
		TitleName: "埃塔尤拉特拉霍",
		TitleCode: "b_idtodyolataloho",
	}
}
