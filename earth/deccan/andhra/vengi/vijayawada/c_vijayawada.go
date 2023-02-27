package vijayawada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VijayawadaCounty interface {
    feud.County
    BChebrolu丘布罗尔() 	feud.Barony
    BGuntur贡土尔() 	feud.Barony
    BKondavidu孔达维杜() 	feud.Barony
    BMasula摩苏拉() 	feud.Barony
    BMonvel慕那吠罗() 	feud.Barony
    BPeddapalli佩达帕尔利() 	feud.Barony
    BVijayawada毗阇耶婆陀() 	feud.Barony
}

type 毗阇耶婆陀VijayawadaCounty struct {
	feud.BaseCounty
	丘布罗尔Chebrolu 	feud.Barony
	贡土尔Guntur 	feud.Barony
	孔达维杜Kondavidu 	feud.Barony
	摩苏拉Masula 	feud.Barony
	慕那吠罗Monvel 	feud.Barony
	佩达帕尔利Peddapalli 	feud.Barony
	毗阇耶婆陀Vijayawada 	feud.Barony
}

func (c *毗阇耶婆陀VijayawadaCounty) BChebrolu丘布罗尔() feud.Barony {
	return c.丘布罗尔Chebrolu
}
    
func (c *毗阇耶婆陀VijayawadaCounty) BGuntur贡土尔() feud.Barony {
	return c.贡土尔Guntur
}
    
func (c *毗阇耶婆陀VijayawadaCounty) BKondavidu孔达维杜() feud.Barony {
	return c.孔达维杜Kondavidu
}
    
func (c *毗阇耶婆陀VijayawadaCounty) BMasula摩苏拉() feud.Barony {
	return c.摩苏拉Masula
}
    
func (c *毗阇耶婆陀VijayawadaCounty) BMonvel慕那吠罗() feud.Barony {
	return c.慕那吠罗Monvel
}
    
func (c *毗阇耶婆陀VijayawadaCounty) BPeddapalli佩达帕尔利() feud.Barony {
	return c.佩达帕尔利Peddapalli
}
    
func (c *毗阇耶婆陀VijayawadaCounty) BVijayawada毗阇耶婆陀() feud.Barony {
	return c.毗阇耶婆陀Vijayawada
}
    
var CVijayawada毗阇耶婆陀 VijayawadaCounty = &毗阇耶婆陀VijayawadaCounty{}

func init() {
	f := CVijayawada毗阇耶婆陀.(*毗阇耶婆陀VijayawadaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1221",
		Title:     "vijayawada",
		TitleName: "毗阇耶婆陀",
		TitleCode: "c_vijayawada",
		Baronies:  map[string]feud.Barony{},
	}

	f.丘布罗尔Chebrolu = BChebrolu丘布罗尔
	f.丘布罗尔Chebrolu.SetParent(f)
	
	f.贡土尔Guntur = BGuntur贡土尔
	f.贡土尔Guntur.SetParent(f)
	
	f.孔达维杜Kondavidu = BKondavidu孔达维杜
	f.孔达维杜Kondavidu.SetParent(f)
	
	f.摩苏拉Masula = BMasula摩苏拉
	f.摩苏拉Masula.SetParent(f)
	
	f.慕那吠罗Monvel = BMonvel慕那吠罗
	f.慕那吠罗Monvel.SetParent(f)
	
	f.佩达帕尔利Peddapalli = BPeddapalli佩达帕尔利
	f.佩达帕尔利Peddapalli.SetParent(f)
	
	f.毗阇耶婆陀Vijayawada = BVijayawada毗阇耶婆陀
	f.毗阇耶婆陀Vijayawada.SetParent(f)
	
}
