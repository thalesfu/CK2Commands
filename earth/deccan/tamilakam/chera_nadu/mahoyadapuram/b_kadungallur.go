package mahoyadapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦顿迦卢尔KadungallurBarony struct {
	feud.BaseBarony
}

var BKadungallur迦顿迦卢尔 feud.Barony = &迦顿迦卢尔KadungallurBarony{}

func init() {
    f := BKadungallur迦顿迦卢尔.(*迦顿迦卢尔KadungallurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kadungallur",
		TitleName: "迦顿迦卢尔",
		TitleCode: "b_kadungallur",
	}
}
