package vidisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VidisaCounty interface {
    feud.County
    BBhojpur蒲阇补罗() 	feud.Barony
    BBijamandal毗阇曼荼罗() 	feud.Barony
    BMangla_devi茫伽罗提毗() 	feud.Barony
    BRaisin罗僧() 	feud.Barony
    BSanchi桑支() 	feud.Barony
    BSironj希罗纳() 	feud.Barony
    BVidisa卑提写() 	feud.Barony
}

type 卑提写VidisaCounty struct {
	feud.BaseCounty
	蒲阇补罗Bhojpur 	feud.Barony
	毗阇曼荼罗Bijamandal 	feud.Barony
	茫伽罗提毗Mangla_devi 	feud.Barony
	罗僧Raisin 	feud.Barony
	桑支Sanchi 	feud.Barony
	希罗纳Sironj 	feud.Barony
	卑提写Vidisa 	feud.Barony
}

func (c *卑提写VidisaCounty) BBhojpur蒲阇补罗() feud.Barony {
	return c.蒲阇补罗Bhojpur
}
    
func (c *卑提写VidisaCounty) BBijamandal毗阇曼荼罗() feud.Barony {
	return c.毗阇曼荼罗Bijamandal
}
    
func (c *卑提写VidisaCounty) BMangla_devi茫伽罗提毗() feud.Barony {
	return c.茫伽罗提毗Mangla_devi
}
    
func (c *卑提写VidisaCounty) BRaisin罗僧() feud.Barony {
	return c.罗僧Raisin
}
    
func (c *卑提写VidisaCounty) BSanchi桑支() feud.Barony {
	return c.桑支Sanchi
}
    
func (c *卑提写VidisaCounty) BSironj希罗纳() feud.Barony {
	return c.希罗纳Sironj
}
    
func (c *卑提写VidisaCounty) BVidisa卑提写() feud.Barony {
	return c.卑提写Vidisa
}
    
var CVidisa卑提写 VidisaCounty = &卑提写VidisaCounty{}

func init() {
	f := CVidisa卑提写.(*卑提写VidisaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1170",
		Title:     "vidisa",
		TitleName: "卑提写",
		TitleCode: "c_vidisa",
		Baronies:  map[string]feud.Barony{},
	}

	f.蒲阇补罗Bhojpur = BBhojpur蒲阇补罗
	f.蒲阇补罗Bhojpur.SetParent(f)
	
	f.毗阇曼荼罗Bijamandal = BBijamandal毗阇曼荼罗
	f.毗阇曼荼罗Bijamandal.SetParent(f)
	
	f.茫伽罗提毗Mangla_devi = BMangla_devi茫伽罗提毗
	f.茫伽罗提毗Mangla_devi.SetParent(f)
	
	f.罗僧Raisin = BRaisin罗僧
	f.罗僧Raisin.SetParent(f)
	
	f.桑支Sanchi = BSanchi桑支
	f.桑支Sanchi.SetParent(f)
	
	f.希罗纳Sironj = BSironj希罗纳
	f.希罗纳Sironj.SetParent(f)
	
	f.卑提写Vidisa = BVidisa卑提写
	f.卑提写Vidisa.SetParent(f)
	
}
