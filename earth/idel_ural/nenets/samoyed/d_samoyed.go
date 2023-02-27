package samoyed

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/samoyed/kanin"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/samoyed/mezen"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/samoyed/samoyeds"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/samoyed/ugra"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SamoyedDuke interface {
    feud.Duke
    CKanin卡宁() 	kanin.KaninCounty
    CMezen梅津() 	mezen.MezenCounty
    CSamoyeds斯诺帕() 	samoyeds.SamoyedsCounty
    CUgra乌格拉() 	ugra.UgraCounty
}

type 梅津SamoyedDuke struct {
	feud.BaseDuke
	卡宁Kanin 	kanin.KaninCounty
	梅津Mezen 	mezen.MezenCounty
	斯诺帕Samoyeds 	samoyeds.SamoyedsCounty
	乌格拉Ugra 	ugra.UgraCounty
}

func (d *梅津SamoyedDuke) CKanin卡宁() kanin.KaninCounty {
	return d.卡宁Kanin
}
    
func (d *梅津SamoyedDuke) CMezen梅津() mezen.MezenCounty {
	return d.梅津Mezen
}
    
func (d *梅津SamoyedDuke) CSamoyeds斯诺帕() samoyeds.SamoyedsCounty {
	return d.斯诺帕Samoyeds
}
    
func (d *梅津SamoyedDuke) CUgra乌格拉() ugra.UgraCounty {
	return d.乌格拉Ugra
}
    
var DSamoyed梅津 SamoyedDuke = &梅津SamoyedDuke{}

func init() {
	f := DSamoyed梅津.(*梅津SamoyedDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "samoyed",
		TitleName: "梅津",
		TitleCode: "d_samoyed",
		Counties:  map[string]feud.County{},
	}

	f.卡宁Kanin = kanin.CKanin卡宁
	f.卡宁Kanin.SetParent(f)
	
	f.梅津Mezen = mezen.CMezen梅津
	f.梅津Mezen.SetParent(f)
	
	f.斯诺帕Samoyeds = samoyeds.CSamoyeds斯诺帕
	f.斯诺帕Samoyeds.SetParent(f)
	
	f.乌格拉Ugra = ugra.CUgra乌格拉
	f.乌格拉Ugra.SetParent(f)
	
}
