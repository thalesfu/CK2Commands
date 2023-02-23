package valencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托伦特TorrentBarony struct {
	feud.BaseBarony
}

var BTorrent托伦特 feud.Barony = &托伦特TorrentBarony{}

func init() {
	f := BTorrent托伦特.(*托伦特TorrentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torrent",
		TitleName: "托伦特",
		TitleCode: "b_torrent",
	}
}
