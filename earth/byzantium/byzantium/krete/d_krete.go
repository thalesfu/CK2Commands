package krete

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/krete/chandax"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/krete/kaneia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KreteDuke interface {
    feud.Duke
    CChandax汉达克斯() 	chandax.ChandaxCounty
    CKaneia卡尼雅() 	kaneia.KaneiaCounty
}

type 克里特KreteDuke struct {
	feud.BaseDuke
	汉达克斯Chandax 	chandax.ChandaxCounty
	卡尼雅Kaneia 	kaneia.KaneiaCounty
}

func (d *克里特KreteDuke) CChandax汉达克斯() chandax.ChandaxCounty {
	return d.汉达克斯Chandax
}
    
func (d *克里特KreteDuke) CKaneia卡尼雅() kaneia.KaneiaCounty {
	return d.卡尼雅Kaneia
}
    
var DKrete克里特 KreteDuke = &克里特KreteDuke{}

func init() {
	f := DKrete克里特.(*克里特KreteDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "krete",
		TitleName: "克里特",
		TitleCode: "d_krete",
		Counties:  map[string]feud.County{},
	}

	f.汉达克斯Chandax = chandax.CChandax汉达克斯
	f.汉达克斯Chandax.SetParent(f)
	
	f.卡尼雅Kaneia = kaneia.CKaneia卡尼雅
	f.卡尼雅Kaneia.SetParent(f)
	
}
