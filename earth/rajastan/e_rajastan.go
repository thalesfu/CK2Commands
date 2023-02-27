package rajastan

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala"
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana"
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RajastanEmpire interface {
    feud.Empire
    KDelhi德里() 	delhi.DelhiKingdom
    KGujarat瞿折罗() 	gujarat.GujaratKingdom
    KKosala憍萨罗() 	kosala.KosalaKingdom
    KMalwa摩腊婆() 	malwa.MalwaKingdom
    KPunjab五河() 	punjab.PunjabKingdom
    KRajputana罗阇弗多那() 	rajputana.RajputanaKingdom
    KSindh信度() 	sindh.SindhKingdom
}

type 罗阇萨傥那RajastanEmpire struct {
	feud.BaseEmpire
	德里Delhi 	delhi.DelhiKingdom
	瞿折罗Gujarat 	gujarat.GujaratKingdom
	憍萨罗Kosala 	kosala.KosalaKingdom
	摩腊婆Malwa 	malwa.MalwaKingdom
	五河Punjab 	punjab.PunjabKingdom
	罗阇弗多那Rajputana 	rajputana.RajputanaKingdom
	信度Sindh 	sindh.SindhKingdom
}

func (e *罗阇萨傥那RajastanEmpire) KDelhi德里() delhi.DelhiKingdom {
	return e.德里Delhi
}
    
func (e *罗阇萨傥那RajastanEmpire) KGujarat瞿折罗() gujarat.GujaratKingdom {
	return e.瞿折罗Gujarat
}
    
func (e *罗阇萨傥那RajastanEmpire) KKosala憍萨罗() kosala.KosalaKingdom {
	return e.憍萨罗Kosala
}
    
func (e *罗阇萨傥那RajastanEmpire) KMalwa摩腊婆() malwa.MalwaKingdom {
	return e.摩腊婆Malwa
}
    
func (e *罗阇萨傥那RajastanEmpire) KPunjab五河() punjab.PunjabKingdom {
	return e.五河Punjab
}
    
func (e *罗阇萨傥那RajastanEmpire) KRajputana罗阇弗多那() rajputana.RajputanaKingdom {
	return e.罗阇弗多那Rajputana
}
    
func (e *罗阇萨傥那RajastanEmpire) KSindh信度() sindh.SindhKingdom {
	return e.信度Sindh
}
    
var ERajastan罗阇萨傥那 RajastanEmpire = &罗阇萨傥那RajastanEmpire{}

func init() {
	f := ERajastan罗阇萨傥那.(*罗阇萨傥那RajastanEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "rajastan",
		TitleName: "罗阇萨傥那",
		TitleCode: "e_rajastan",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.德里Delhi = delhi.KDelhi德里
	f.德里Delhi.SetParent(f)
	f.瞿折罗Gujarat = gujarat.KGujarat瞿折罗
	f.瞿折罗Gujarat.SetParent(f)
	f.憍萨罗Kosala = kosala.KKosala憍萨罗
	f.憍萨罗Kosala.SetParent(f)
	f.摩腊婆Malwa = malwa.KMalwa摩腊婆
	f.摩腊婆Malwa.SetParent(f)
	f.五河Punjab = punjab.KPunjab五河
	f.五河Punjab.SetParent(f)
	f.罗阇弗多那Rajputana = rajputana.KRajputana罗阇弗多那
	f.罗阇弗多那Rajputana.SetParent(f)
	f.信度Sindh = sindh.KSindh信度
	f.信度Sindh.SetParent(f)
}
