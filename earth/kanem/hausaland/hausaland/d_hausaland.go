package hausaland

import (
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/hausaland/daura"
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/hausaland/gobir"
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/hausaland/kano"
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/hausaland/katsina"
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/hausaland/zamfara"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HausalandDuke interface {
    feud.Duke
    CDaura道拉() 	daura.DauraCounty
    CGobir戈比尔() 	gobir.GobirCounty
    CKano卡诺() 	kano.KanoCounty
    CKatsina卡齐纳() 	katsina.KatsinaCounty
    CZamfara扎姆法拉() 	zamfara.ZamfaraCounty
}

type 豪萨HausalandDuke struct {
	feud.BaseDuke
	道拉Daura 	daura.DauraCounty
	戈比尔Gobir 	gobir.GobirCounty
	卡诺Kano 	kano.KanoCounty
	卡齐纳Katsina 	katsina.KatsinaCounty
	扎姆法拉Zamfara 	zamfara.ZamfaraCounty
}

func (d *豪萨HausalandDuke) CDaura道拉() daura.DauraCounty {
	return d.道拉Daura
}
    
func (d *豪萨HausalandDuke) CGobir戈比尔() gobir.GobirCounty {
	return d.戈比尔Gobir
}
    
func (d *豪萨HausalandDuke) CKano卡诺() kano.KanoCounty {
	return d.卡诺Kano
}
    
func (d *豪萨HausalandDuke) CKatsina卡齐纳() katsina.KatsinaCounty {
	return d.卡齐纳Katsina
}
    
func (d *豪萨HausalandDuke) CZamfara扎姆法拉() zamfara.ZamfaraCounty {
	return d.扎姆法拉Zamfara
}
    
var DHausaland豪萨 HausalandDuke = &豪萨HausalandDuke{}

func init() {
	f := DHausaland豪萨.(*豪萨HausalandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "hausaland",
		TitleName: "豪萨",
		TitleCode: "d_hausaland",
		Counties:  map[string]feud.County{},
	}

	f.道拉Daura = daura.CDaura道拉
	f.道拉Daura.SetParent(f)
	
	f.戈比尔Gobir = gobir.CGobir戈比尔
	f.戈比尔Gobir.SetParent(f)
	
	f.卡诺Kano = kano.CKano卡诺
	f.卡诺Kano.SetParent(f)
	
	f.卡齐纳Katsina = katsina.CKatsina卡齐纳
	f.卡齐纳Katsina.SetParent(f)
	
	f.扎姆法拉Zamfara = zamfara.CZamfara扎姆法拉
	f.扎姆法拉Zamfara.SetParent(f)
	
}
