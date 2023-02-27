package yazd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔夫特TaftBarony struct {
	feud.BaseBarony
}

var BTaft塔夫特 feud.Barony = &塔夫特TaftBarony{}

func init() {
    f := BTaft塔夫特.(*塔夫特TaftBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taft",
		TitleName: "塔夫特",
		TitleCode: "b_taft",
	}
}
