package udine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔琴托FriuliBarony struct {
	feud.BaseBarony
}

var BFriuli塔尔琴托 feud.Barony = &塔尔琴托FriuliBarony{}

func init() {
	f := BFriuli塔尔琴托.(*塔尔琴托FriuliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "friuli",
		TitleName: "塔尔琴托",
		TitleCode: "b_friuli",
	}
}
