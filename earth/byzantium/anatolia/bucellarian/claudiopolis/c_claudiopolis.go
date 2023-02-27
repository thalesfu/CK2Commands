package claudiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ClaudiopolisCounty interface {
    feud.County
    BBithynium比提尼乌姆() 	feud.Barony
    BClaudiopolis克劳狄奥波利斯() 	feud.Barony
    BHonorias奥诺里亚斯() 	feud.Barony
    BIuliopolis尤利奥波利斯() 	feud.Barony
    BPessinus培希努() 	feud.Barony
    BPetobriga派特布里格() 	feud.Barony
    BSangarius桑加里奥斯() 	feud.Barony
}

type 克劳狄奥波利斯ClaudiopolisCounty struct {
	feud.BaseCounty
	比提尼乌姆Bithynium 	feud.Barony
	克劳狄奥波利斯Claudiopolis 	feud.Barony
	奥诺里亚斯Honorias 	feud.Barony
	尤利奥波利斯Iuliopolis 	feud.Barony
	培希努Pessinus 	feud.Barony
	派特布里格Petobriga 	feud.Barony
	桑加里奥斯Sangarius 	feud.Barony
}

func (c *克劳狄奥波利斯ClaudiopolisCounty) BBithynium比提尼乌姆() feud.Barony {
	return c.比提尼乌姆Bithynium
}
    
func (c *克劳狄奥波利斯ClaudiopolisCounty) BClaudiopolis克劳狄奥波利斯() feud.Barony {
	return c.克劳狄奥波利斯Claudiopolis
}
    
func (c *克劳狄奥波利斯ClaudiopolisCounty) BHonorias奥诺里亚斯() feud.Barony {
	return c.奥诺里亚斯Honorias
}
    
func (c *克劳狄奥波利斯ClaudiopolisCounty) BIuliopolis尤利奥波利斯() feud.Barony {
	return c.尤利奥波利斯Iuliopolis
}
    
func (c *克劳狄奥波利斯ClaudiopolisCounty) BPessinus培希努() feud.Barony {
	return c.培希努Pessinus
}
    
func (c *克劳狄奥波利斯ClaudiopolisCounty) BPetobriga派特布里格() feud.Barony {
	return c.派特布里格Petobriga
}
    
func (c *克劳狄奥波利斯ClaudiopolisCounty) BSangarius桑加里奥斯() feud.Barony {
	return c.桑加里奥斯Sangarius
}
    
var CClaudiopolis克劳狄奥波利斯 ClaudiopolisCounty = &克劳狄奥波利斯ClaudiopolisCounty{}

func init() {
	f := CClaudiopolis克劳狄奥波利斯.(*克劳狄奥波利斯ClaudiopolisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1930",
		Title:     "claudiopolis",
		TitleName: "克劳狄奥波利斯",
		TitleCode: "c_claudiopolis",
		Baronies:  map[string]feud.Barony{},
	}

	f.比提尼乌姆Bithynium = BBithynium比提尼乌姆
	f.比提尼乌姆Bithynium.SetParent(f)
	
	f.克劳狄奥波利斯Claudiopolis = BClaudiopolis克劳狄奥波利斯
	f.克劳狄奥波利斯Claudiopolis.SetParent(f)
	
	f.奥诺里亚斯Honorias = BHonorias奥诺里亚斯
	f.奥诺里亚斯Honorias.SetParent(f)
	
	f.尤利奥波利斯Iuliopolis = BIuliopolis尤利奥波利斯
	f.尤利奥波利斯Iuliopolis.SetParent(f)
	
	f.培希努Pessinus = BPessinus培希努
	f.培希努Pessinus.SetParent(f)
	
	f.派特布里格Petobriga = BPetobriga派特布里格
	f.派特布里格Petobriga.SetParent(f)
	
	f.桑加里奥斯Sangarius = BSangarius桑加里奥斯
	f.桑加里奥斯Sangarius.SetParent(f)
	
}
