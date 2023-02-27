package burhanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BurhanpurCounty interface {
    feud.County
    BBurhanpur补罗汉补罗() 	feud.Barony
    BChangdev章伽提婆() 	feud.Barony
    BErandol埃伦多尔() 	feud.Barony
    BHalgeri诃罗祇梨() 	feud.Barony
    BJargaon折罗伽罗摩() 	feud.Barony
    BVaghli伐迦利() 	feud.Barony
    BYawal耶瓦() 	feud.Barony
}

type 补罗汉补罗BurhanpurCounty struct {
	feud.BaseCounty
	补罗汉补罗Burhanpur 	feud.Barony
	章伽提婆Changdev 	feud.Barony
	埃伦多尔Erandol 	feud.Barony
	诃罗祇梨Halgeri 	feud.Barony
	折罗伽罗摩Jargaon 	feud.Barony
	伐迦利Vaghli 	feud.Barony
	耶瓦Yawal 	feud.Barony
}

func (c *补罗汉补罗BurhanpurCounty) BBurhanpur补罗汉补罗() feud.Barony {
	return c.补罗汉补罗Burhanpur
}
    
func (c *补罗汉补罗BurhanpurCounty) BChangdev章伽提婆() feud.Barony {
	return c.章伽提婆Changdev
}
    
func (c *补罗汉补罗BurhanpurCounty) BErandol埃伦多尔() feud.Barony {
	return c.埃伦多尔Erandol
}
    
func (c *补罗汉补罗BurhanpurCounty) BHalgeri诃罗祇梨() feud.Barony {
	return c.诃罗祇梨Halgeri
}
    
func (c *补罗汉补罗BurhanpurCounty) BJargaon折罗伽罗摩() feud.Barony {
	return c.折罗伽罗摩Jargaon
}
    
func (c *补罗汉补罗BurhanpurCounty) BVaghli伐迦利() feud.Barony {
	return c.伐迦利Vaghli
}
    
func (c *补罗汉补罗BurhanpurCounty) BYawal耶瓦() feud.Barony {
	return c.耶瓦Yawal
}
    
var CBurhanpur补罗汉补罗 BurhanpurCounty = &补罗汉补罗BurhanpurCounty{}

func init() {
	f := CBurhanpur补罗汉补罗.(*补罗汉补罗BurhanpurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1263",
		Title:     "burhanpur",
		TitleName: "补罗汉补罗",
		TitleCode: "c_burhanpur",
		Baronies:  map[string]feud.Barony{},
	}

	f.补罗汉补罗Burhanpur = BBurhanpur补罗汉补罗
	f.补罗汉补罗Burhanpur.SetParent(f)
	
	f.章伽提婆Changdev = BChangdev章伽提婆
	f.章伽提婆Changdev.SetParent(f)
	
	f.埃伦多尔Erandol = BErandol埃伦多尔
	f.埃伦多尔Erandol.SetParent(f)
	
	f.诃罗祇梨Halgeri = BHalgeri诃罗祇梨
	f.诃罗祇梨Halgeri.SetParent(f)
	
	f.折罗伽罗摩Jargaon = BJargaon折罗伽罗摩
	f.折罗伽罗摩Jargaon.SetParent(f)
	
	f.伐迦利Vaghli = BVaghli伐迦利
	f.伐迦利Vaghli.SetParent(f)
	
	f.耶瓦Yawal = BYawal耶瓦
	f.耶瓦Yawal.SetParent(f)
	
}
