package thracesia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/thracesia/laodikeia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/thracesia/sardis"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/thracesia/sozopolis"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThracesiaDuke interface {
    feud.Duke
    CLaodikeia老底嘉() 	laodikeia.LaodikeiaCounty
    CSardis萨第斯() 	sardis.SardisCounty
    CSozopolis索佐波利斯() 	sozopolis.SozopolisCounty
}

type 色雷斯西亚ThracesiaDuke struct {
	feud.BaseDuke
	老底嘉Laodikeia 	laodikeia.LaodikeiaCounty
	萨第斯Sardis 	sardis.SardisCounty
	索佐波利斯Sozopolis 	sozopolis.SozopolisCounty
}

func (d *色雷斯西亚ThracesiaDuke) CLaodikeia老底嘉() laodikeia.LaodikeiaCounty {
	return d.老底嘉Laodikeia
}
    
func (d *色雷斯西亚ThracesiaDuke) CSardis萨第斯() sardis.SardisCounty {
	return d.萨第斯Sardis
}
    
func (d *色雷斯西亚ThracesiaDuke) CSozopolis索佐波利斯() sozopolis.SozopolisCounty {
	return d.索佐波利斯Sozopolis
}
    
var DThracesia色雷斯西亚 ThracesiaDuke = &色雷斯西亚ThracesiaDuke{}

func init() {
	f := DThracesia色雷斯西亚.(*色雷斯西亚ThracesiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "thracesia",
		TitleName: "色雷斯西亚",
		TitleCode: "d_thracesia",
		Counties:  map[string]feud.County{},
	}

	f.老底嘉Laodikeia = laodikeia.CLaodikeia老底嘉
	f.老底嘉Laodikeia.SetParent(f)
	
	f.萨第斯Sardis = sardis.CSardis萨第斯
	f.萨第斯Sardis.SetParent(f)
	
	f.索佐波利斯Sozopolis = sozopolis.CSozopolis索佐波利斯
	f.索佐波利斯Sozopolis.SetParent(f)
	
}
