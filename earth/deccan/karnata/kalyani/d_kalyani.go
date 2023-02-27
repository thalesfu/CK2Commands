package kalyani

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/kalyani/bidar"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/kalyani/kalyani"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/kalyani/manyakheta"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/kalyani/sagar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KalyaniDuke interface {
    feud.Duke
    CBidar毗陀罗() 	bidar.BidarCounty
    CKalyani迦梨耶尼() 	kalyani.KalyaniCounty
    CManyakheta摩若契吒() 	manyakheta.ManyakhetaCounty
    CSagar娑伽罗() 	sagar.SagarCounty
}

type 迦梨耶尼KalyaniDuke struct {
	feud.BaseDuke
	毗陀罗Bidar 	bidar.BidarCounty
	迦梨耶尼Kalyani 	kalyani.KalyaniCounty
	摩若契吒Manyakheta 	manyakheta.ManyakhetaCounty
	娑伽罗Sagar 	sagar.SagarCounty
}

func (d *迦梨耶尼KalyaniDuke) CBidar毗陀罗() bidar.BidarCounty {
	return d.毗陀罗Bidar
}
    
func (d *迦梨耶尼KalyaniDuke) CKalyani迦梨耶尼() kalyani.KalyaniCounty {
	return d.迦梨耶尼Kalyani
}
    
func (d *迦梨耶尼KalyaniDuke) CManyakheta摩若契吒() manyakheta.ManyakhetaCounty {
	return d.摩若契吒Manyakheta
}
    
func (d *迦梨耶尼KalyaniDuke) CSagar娑伽罗() sagar.SagarCounty {
	return d.娑伽罗Sagar
}
    
var DKalyani迦梨耶尼 KalyaniDuke = &迦梨耶尼KalyaniDuke{}

func init() {
	f := DKalyani迦梨耶尼.(*迦梨耶尼KalyaniDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kalyani",
		TitleName: "迦梨耶尼",
		TitleCode: "d_kalyani",
		Counties:  map[string]feud.County{},
	}

	f.毗陀罗Bidar = bidar.CBidar毗陀罗
	f.毗陀罗Bidar.SetParent(f)
	
	f.迦梨耶尼Kalyani = kalyani.CKalyani迦梨耶尼
	f.迦梨耶尼Kalyani.SetParent(f)
	
	f.摩若契吒Manyakheta = manyakheta.CManyakheta摩若契吒
	f.摩若契吒Manyakheta.SetParent(f)
	
	f.娑伽罗Sagar = sagar.CSagar娑伽罗
	f.娑伽罗Sagar.SetParent(f)
	
}
