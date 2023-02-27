package paphlagonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PaphlagoniaCounty interface {
    feud.County
    BAnastasiopolis阿纳塔修波利斯() 	feud.Barony
    BBolu博卢() 	feud.Barony
    BCabira加比拉() 	feud.Barony
    BGangra冈格拉() 	feud.Barony
    BLeontopolis莱翁托波利斯() 	feud.Barony
    BSafranbolu萨夫兰博卢() 	feud.Barony
    BZaliscus匝利刻斯() 	feud.Barony
}

type 帕夫拉戈尼亚PaphlagoniaCounty struct {
	feud.BaseCounty
	阿纳塔修波利斯Anastasiopolis 	feud.Barony
	博卢Bolu 	feud.Barony
	加比拉Cabira 	feud.Barony
	冈格拉Gangra 	feud.Barony
	莱翁托波利斯Leontopolis 	feud.Barony
	萨夫兰博卢Safranbolu 	feud.Barony
	匝利刻斯Zaliscus 	feud.Barony
}

func (c *帕夫拉戈尼亚PaphlagoniaCounty) BAnastasiopolis阿纳塔修波利斯() feud.Barony {
	return c.阿纳塔修波利斯Anastasiopolis
}
    
func (c *帕夫拉戈尼亚PaphlagoniaCounty) BBolu博卢() feud.Barony {
	return c.博卢Bolu
}
    
func (c *帕夫拉戈尼亚PaphlagoniaCounty) BCabira加比拉() feud.Barony {
	return c.加比拉Cabira
}
    
func (c *帕夫拉戈尼亚PaphlagoniaCounty) BGangra冈格拉() feud.Barony {
	return c.冈格拉Gangra
}
    
func (c *帕夫拉戈尼亚PaphlagoniaCounty) BLeontopolis莱翁托波利斯() feud.Barony {
	return c.莱翁托波利斯Leontopolis
}
    
func (c *帕夫拉戈尼亚PaphlagoniaCounty) BSafranbolu萨夫兰博卢() feud.Barony {
	return c.萨夫兰博卢Safranbolu
}
    
func (c *帕夫拉戈尼亚PaphlagoniaCounty) BZaliscus匝利刻斯() feud.Barony {
	return c.匝利刻斯Zaliscus
}
    
var CPaphlagonia帕夫拉戈尼亚 PaphlagoniaCounty = &帕夫拉戈尼亚PaphlagoniaCounty{}

func init() {
	f := CPaphlagonia帕夫拉戈尼亚.(*帕夫拉戈尼亚PaphlagoniaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "751",
		Title:     "paphlagonia",
		TitleName: "帕夫拉戈尼亚",
		TitleCode: "c_paphlagonia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿纳塔修波利斯Anastasiopolis = BAnastasiopolis阿纳塔修波利斯
	f.阿纳塔修波利斯Anastasiopolis.SetParent(f)
	
	f.博卢Bolu = BBolu博卢
	f.博卢Bolu.SetParent(f)
	
	f.加比拉Cabira = BCabira加比拉
	f.加比拉Cabira.SetParent(f)
	
	f.冈格拉Gangra = BGangra冈格拉
	f.冈格拉Gangra.SetParent(f)
	
	f.莱翁托波利斯Leontopolis = BLeontopolis莱翁托波利斯
	f.莱翁托波利斯Leontopolis.SetParent(f)
	
	f.萨夫兰博卢Safranbolu = BSafranbolu萨夫兰博卢
	f.萨夫兰博卢Safranbolu.SetParent(f)
	
	f.匝利刻斯Zaliscus = BZaliscus匝利刻斯
	f.匝利刻斯Zaliscus.SetParent(f)
	
}
