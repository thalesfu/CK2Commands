package bedford

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/bedford/bedford"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/bedford/essex"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/bedford/middlesex"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/bedford/northampton"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BedfordDuke interface {
    feud.Duke
    CBedford贝德福德() 	bedford.BedfordCounty
    CEssex埃塞克斯() 	essex.EssexCounty
    CMiddlesex米德尔塞克斯() 	middlesex.MiddlesexCounty
    CNorthampton北安普顿() 	northampton.NorthamptonCounty
}

type 埃塞克斯BedfordDuke struct {
	feud.BaseDuke
	贝德福德Bedford 	bedford.BedfordCounty
	埃塞克斯Essex 	essex.EssexCounty
	米德尔塞克斯Middlesex 	middlesex.MiddlesexCounty
	北安普顿Northampton 	northampton.NorthamptonCounty
}

func (d *埃塞克斯BedfordDuke) CBedford贝德福德() bedford.BedfordCounty {
	return d.贝德福德Bedford
}
    
func (d *埃塞克斯BedfordDuke) CEssex埃塞克斯() essex.EssexCounty {
	return d.埃塞克斯Essex
}
    
func (d *埃塞克斯BedfordDuke) CMiddlesex米德尔塞克斯() middlesex.MiddlesexCounty {
	return d.米德尔塞克斯Middlesex
}
    
func (d *埃塞克斯BedfordDuke) CNorthampton北安普顿() northampton.NorthamptonCounty {
	return d.北安普顿Northampton
}
    
var DBedford埃塞克斯 BedfordDuke = &埃塞克斯BedfordDuke{}

func init() {
	f := DBedford埃塞克斯.(*埃塞克斯BedfordDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bedford",
		TitleName: "埃塞克斯",
		TitleCode: "d_bedford",
		Counties:  map[string]feud.County{},
	}

	f.贝德福德Bedford = bedford.CBedford贝德福德
	f.贝德福德Bedford.SetParent(f)
	
	f.埃塞克斯Essex = essex.CEssex埃塞克斯
	f.埃塞克斯Essex.SetParent(f)
	
	f.米德尔塞克斯Middlesex = middlesex.CMiddlesex米德尔塞克斯
	f.米德尔塞克斯Middlesex.SetParent(f)
	
	f.北安普顿Northampton = northampton.CNorthampton北安普顿
	f.北安普顿Northampton.SetParent(f)
	
}
