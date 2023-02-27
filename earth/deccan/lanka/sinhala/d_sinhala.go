package sinhala

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/lanka/sinhala/kotthasara"
	"github.com/thalesfu/CK2Commands/earth/deccan/lanka/sinhala/nagadipa"
	"github.com/thalesfu/CK2Commands/earth/deccan/lanka/sinhala/phiti"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SinhalaDuke interface {
    feud.Duke
    CKotthasara拘吒娑罗() 	kotthasara.KotthasaraCounty
    CNagadipa那伽洲() 	nagadipa.NagadipaCounty
    CPhiti比醯帝() 	phiti.PhitiCounty
}

type 僧伽罗SinhalaDuke struct {
	feud.BaseDuke
	拘吒娑罗Kotthasara 	kotthasara.KotthasaraCounty
	那伽洲Nagadipa 	nagadipa.NagadipaCounty
	比醯帝Phiti 	phiti.PhitiCounty
}

func (d *僧伽罗SinhalaDuke) CKotthasara拘吒娑罗() kotthasara.KotthasaraCounty {
	return d.拘吒娑罗Kotthasara
}
    
func (d *僧伽罗SinhalaDuke) CNagadipa那伽洲() nagadipa.NagadipaCounty {
	return d.那伽洲Nagadipa
}
    
func (d *僧伽罗SinhalaDuke) CPhiti比醯帝() phiti.PhitiCounty {
	return d.比醯帝Phiti
}
    
var DSinhala僧伽罗 SinhalaDuke = &僧伽罗SinhalaDuke{}

func init() {
	f := DSinhala僧伽罗.(*僧伽罗SinhalaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sinhala",
		TitleName: "僧伽罗",
		TitleCode: "d_sinhala",
		Counties:  map[string]feud.County{},
	}

	f.拘吒娑罗Kotthasara = kotthasara.CKotthasara拘吒娑罗
	f.拘吒娑罗Kotthasara.SetParent(f)
	
	f.那伽洲Nagadipa = nagadipa.CNagadipa那伽洲
	f.那伽洲Nagadipa.SetParent(f)
	
	f.比醯帝Phiti = phiti.CPhiti比醯帝
	f.比醯帝Phiti.SetParent(f)
	
}
