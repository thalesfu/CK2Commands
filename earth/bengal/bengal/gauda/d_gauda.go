package gauda

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/gauda/gauda"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/gauda/kotivarsa"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/gauda/laksmanavati"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/gauda/radha"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GaudaDuke interface {
    feud.Duke
    CGauda乔荼() 	gauda.GaudaCounty
    CKotivarsa俱知勃里沙() 	kotivarsa.KotivarsaCounty
    CLaksmanavati罗叉摩那伐底() 	laksmanavati.LaksmanavatiCounty
    CRadha罗陀() 	radha.RadhaCounty
}

type 乔荼GaudaDuke struct {
	feud.BaseDuke
	乔荼Gauda 	gauda.GaudaCounty
	俱知勃里沙Kotivarsa 	kotivarsa.KotivarsaCounty
	罗叉摩那伐底Laksmanavati 	laksmanavati.LaksmanavatiCounty
	罗陀Radha 	radha.RadhaCounty
}

func (d *乔荼GaudaDuke) CGauda乔荼() gauda.GaudaCounty {
	return d.乔荼Gauda
}
    
func (d *乔荼GaudaDuke) CKotivarsa俱知勃里沙() kotivarsa.KotivarsaCounty {
	return d.俱知勃里沙Kotivarsa
}
    
func (d *乔荼GaudaDuke) CLaksmanavati罗叉摩那伐底() laksmanavati.LaksmanavatiCounty {
	return d.罗叉摩那伐底Laksmanavati
}
    
func (d *乔荼GaudaDuke) CRadha罗陀() radha.RadhaCounty {
	return d.罗陀Radha
}
    
var DGauda乔荼 GaudaDuke = &乔荼GaudaDuke{}

func init() {
	f := DGauda乔荼.(*乔荼GaudaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gauda",
		TitleName: "乔荼",
		TitleCode: "d_gauda",
		Counties:  map[string]feud.County{},
	}

	f.乔荼Gauda = gauda.CGauda乔荼
	f.乔荼Gauda.SetParent(f)
	
	f.俱知勃里沙Kotivarsa = kotivarsa.CKotivarsa俱知勃里沙
	f.俱知勃里沙Kotivarsa.SetParent(f)
	
	f.罗叉摩那伐底Laksmanavati = laksmanavati.CLaksmanavati罗叉摩那伐底
	f.罗叉摩那伐底Laksmanavati.SetParent(f)
	
	f.罗陀Radha = radha.CRadha罗陀
	f.罗陀Radha.SetParent(f)
	
}
