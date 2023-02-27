package gangavadi

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/gangavadi/manyapura"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/gangavadi/nandagiri"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/gangavadi/srirangapatna"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/gangavadi/talakad"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GangavadiDuke interface {
    feud.Duke
    CManyapura摩尼耶补罗() 	manyapura.ManyapuraCounty
    CNandagiri难陀山() 	nandagiri.NandagiriCounty
    CSrirangapatna室利浪伽钵多那() 	srirangapatna.SrirangapatnaCounty
    CTalakad多罗迦都() 	talakad.TalakadCounty
}

type 恒伽婆稚GangavadiDuke struct {
	feud.BaseDuke
	摩尼耶补罗Manyapura 	manyapura.ManyapuraCounty
	难陀山Nandagiri 	nandagiri.NandagiriCounty
	室利浪伽钵多那Srirangapatna 	srirangapatna.SrirangapatnaCounty
	多罗迦都Talakad 	talakad.TalakadCounty
}

func (d *恒伽婆稚GangavadiDuke) CManyapura摩尼耶补罗() manyapura.ManyapuraCounty {
	return d.摩尼耶补罗Manyapura
}
    
func (d *恒伽婆稚GangavadiDuke) CNandagiri难陀山() nandagiri.NandagiriCounty {
	return d.难陀山Nandagiri
}
    
func (d *恒伽婆稚GangavadiDuke) CSrirangapatna室利浪伽钵多那() srirangapatna.SrirangapatnaCounty {
	return d.室利浪伽钵多那Srirangapatna
}
    
func (d *恒伽婆稚GangavadiDuke) CTalakad多罗迦都() talakad.TalakadCounty {
	return d.多罗迦都Talakad
}
    
var DGangavadi恒伽婆稚 GangavadiDuke = &恒伽婆稚GangavadiDuke{}

func init() {
	f := DGangavadi恒伽婆稚.(*恒伽婆稚GangavadiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gangavadi",
		TitleName: "恒伽婆稚",
		TitleCode: "d_gangavadi",
		Counties:  map[string]feud.County{},
	}

	f.摩尼耶补罗Manyapura = manyapura.CManyapura摩尼耶补罗
	f.摩尼耶补罗Manyapura.SetParent(f)
	
	f.难陀山Nandagiri = nandagiri.CNandagiri难陀山
	f.难陀山Nandagiri.SetParent(f)
	
	f.室利浪伽钵多那Srirangapatna = srirangapatna.CSrirangapatna室利浪伽钵多那
	f.室利浪伽钵多那Srirangapatna.SetParent(f)
	
	f.多罗迦都Talakad = talakad.CTalakad多罗迦都
	f.多罗迦都Talakad.SetParent(f)
	
}
