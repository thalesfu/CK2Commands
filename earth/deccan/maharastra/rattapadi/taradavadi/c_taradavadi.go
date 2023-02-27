package taradavadi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TaradavadiCounty interface {
    feud.County
    BHastikundi诃悉帝军持() 	feud.Barony
    BKarambakkudi伽蓝婆俱提() 	feud.Barony
    BPandharpur般荼罗补罗() 	feud.Barony
    BSangole桑格奥莱() 	feud.Barony
    BSaundatti苏犍陀伐底() 	feud.Barony
    BTaradavadi多罗陀婆稚() 	feud.Barony
    BViyapura毗耶补罗() 	feud.Barony
}

type 多罗陀婆稚TaradavadiCounty struct {
	feud.BaseCounty
	诃悉帝军持Hastikundi 	feud.Barony
	伽蓝婆俱提Karambakkudi 	feud.Barony
	般荼罗补罗Pandharpur 	feud.Barony
	桑格奥莱Sangole 	feud.Barony
	苏犍陀伐底Saundatti 	feud.Barony
	多罗陀婆稚Taradavadi 	feud.Barony
	毗耶补罗Viyapura 	feud.Barony
}

func (c *多罗陀婆稚TaradavadiCounty) BHastikundi诃悉帝军持() feud.Barony {
	return c.诃悉帝军持Hastikundi
}
    
func (c *多罗陀婆稚TaradavadiCounty) BKarambakkudi伽蓝婆俱提() feud.Barony {
	return c.伽蓝婆俱提Karambakkudi
}
    
func (c *多罗陀婆稚TaradavadiCounty) BPandharpur般荼罗补罗() feud.Barony {
	return c.般荼罗补罗Pandharpur
}
    
func (c *多罗陀婆稚TaradavadiCounty) BSangole桑格奥莱() feud.Barony {
	return c.桑格奥莱Sangole
}
    
func (c *多罗陀婆稚TaradavadiCounty) BSaundatti苏犍陀伐底() feud.Barony {
	return c.苏犍陀伐底Saundatti
}
    
func (c *多罗陀婆稚TaradavadiCounty) BTaradavadi多罗陀婆稚() feud.Barony {
	return c.多罗陀婆稚Taradavadi
}
    
func (c *多罗陀婆稚TaradavadiCounty) BViyapura毗耶补罗() feud.Barony {
	return c.毗耶补罗Viyapura
}
    
var CTaradavadi多罗陀婆稚 TaradavadiCounty = &多罗陀婆稚TaradavadiCounty{}

func init() {
	f := CTaradavadi多罗陀婆稚.(*多罗陀婆稚TaradavadiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1206",
		Title:     "taradavadi",
		TitleName: "多罗陀婆稚",
		TitleCode: "c_taradavadi",
		Baronies:  map[string]feud.Barony{},
	}

	f.诃悉帝军持Hastikundi = BHastikundi诃悉帝军持
	f.诃悉帝军持Hastikundi.SetParent(f)
	
	f.伽蓝婆俱提Karambakkudi = BKarambakkudi伽蓝婆俱提
	f.伽蓝婆俱提Karambakkudi.SetParent(f)
	
	f.般荼罗补罗Pandharpur = BPandharpur般荼罗补罗
	f.般荼罗补罗Pandharpur.SetParent(f)
	
	f.桑格奥莱Sangole = BSangole桑格奥莱
	f.桑格奥莱Sangole.SetParent(f)
	
	f.苏犍陀伐底Saundatti = BSaundatti苏犍陀伐底
	f.苏犍陀伐底Saundatti.SetParent(f)
	
	f.多罗陀婆稚Taradavadi = BTaradavadi多罗陀婆稚
	f.多罗陀婆稚Taradavadi.SetParent(f)
	
	f.毗耶补罗Viyapura = BViyapura毗耶补罗
	f.毗耶补罗Viyapura.SetParent(f)
	
}
