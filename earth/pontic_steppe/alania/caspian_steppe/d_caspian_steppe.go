package caspian_steppe

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/caspian_steppe/kalaus"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/caspian_steppe/kuma"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/caspian_steppe/manych"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/caspian_steppe/terek"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Caspian_steppeDuke interface {
    feud.Duke
    CKalaus卡劳斯() 	kalaus.KalausCounty
    CKuma库马() 	kuma.KumaCounty
    CManych马内奇() 	manych.ManychCounty
    CTerek捷列克() 	terek.TerekCounty
}

type 里海草原Caspian_steppeDuke struct {
	feud.BaseDuke
	卡劳斯Kalaus 	kalaus.KalausCounty
	库马Kuma 	kuma.KumaCounty
	马内奇Manych 	manych.ManychCounty
	捷列克Terek 	terek.TerekCounty
}

func (d *里海草原Caspian_steppeDuke) CKalaus卡劳斯() kalaus.KalausCounty {
	return d.卡劳斯Kalaus
}
    
func (d *里海草原Caspian_steppeDuke) CKuma库马() kuma.KumaCounty {
	return d.库马Kuma
}
    
func (d *里海草原Caspian_steppeDuke) CManych马内奇() manych.ManychCounty {
	return d.马内奇Manych
}
    
func (d *里海草原Caspian_steppeDuke) CTerek捷列克() terek.TerekCounty {
	return d.捷列克Terek
}
    
var DCaspian_steppe里海草原 Caspian_steppeDuke = &里海草原Caspian_steppeDuke{}

func init() {
	f := DCaspian_steppe里海草原.(*里海草原Caspian_steppeDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "caspian_steppe",
		TitleName: "里海草原",
		TitleCode: "d_caspian_steppe",
		Counties:  map[string]feud.County{},
	}

	f.卡劳斯Kalaus = kalaus.CKalaus卡劳斯
	f.卡劳斯Kalaus.SetParent(f)
	
	f.库马Kuma = kuma.CKuma库马
	f.库马Kuma.SetParent(f)
	
	f.马内奇Manych = manych.CManych马内奇
	f.马内奇Manych.SetParent(f)
	
	f.捷列克Terek = terek.CTerek捷列克
	f.捷列克Terek.SetParent(f)
	
}
