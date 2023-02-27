package bikrampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗讫罗摩补罗BikrampurBarony struct {
	feud.BaseBarony
}

var BBikrampur毗讫罗摩补罗 feud.Barony = &毗讫罗摩补罗BikrampurBarony{}

func init() {
    f := BBikrampur毗讫罗摩补罗.(*毗讫罗摩补罗BikrampurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bikrampur",
		TitleName: "毗讫罗摩补罗",
		TitleCode: "b_bikrampur",
	}
}
