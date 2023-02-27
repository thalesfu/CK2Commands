package transylvania

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/transylvania/bihar"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/transylvania/feher"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/transylvania/szekelyfold"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TransylvaniaDuke interface {
    feud.Duke
    CBihar比豪尔() 	bihar.BiharCounty
    CFeher费黑尔() 	feher.FeherCounty
    CSzekelyfold塞凯伊弗尔德() 	szekelyfold.SzekelyfoldCounty
}

type 特兰西瓦尼亚TransylvaniaDuke struct {
	feud.BaseDuke
	比豪尔Bihar 	bihar.BiharCounty
	费黑尔Feher 	feher.FeherCounty
	塞凯伊弗尔德Szekelyfold 	szekelyfold.SzekelyfoldCounty
}

func (d *特兰西瓦尼亚TransylvaniaDuke) CBihar比豪尔() bihar.BiharCounty {
	return d.比豪尔Bihar
}
    
func (d *特兰西瓦尼亚TransylvaniaDuke) CFeher费黑尔() feher.FeherCounty {
	return d.费黑尔Feher
}
    
func (d *特兰西瓦尼亚TransylvaniaDuke) CSzekelyfold塞凯伊弗尔德() szekelyfold.SzekelyfoldCounty {
	return d.塞凯伊弗尔德Szekelyfold
}
    
var DTransylvania特兰西瓦尼亚 TransylvaniaDuke = &特兰西瓦尼亚TransylvaniaDuke{}

func init() {
	f := DTransylvania特兰西瓦尼亚.(*特兰西瓦尼亚TransylvaniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "transylvania",
		TitleName: "特兰西瓦尼亚",
		TitleCode: "d_transylvania",
		Counties:  map[string]feud.County{},
	}

	f.比豪尔Bihar = bihar.CBihar比豪尔
	f.比豪尔Bihar.SetParent(f)
	
	f.费黑尔Feher = feher.CFeher费黑尔
	f.费黑尔Feher.SetParent(f)
	
	f.塞凯伊弗尔德Szekelyfold = szekelyfold.CSzekelyfold塞凯伊弗尔德
	f.塞凯伊弗尔德Szekelyfold.SetParent(f)
	
}
