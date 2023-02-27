package chera_nadu

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/chera_nadu/kanara"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/chera_nadu/mahoyadapuram"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/chera_nadu/maldives"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/chera_nadu/qalqut"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/chera_nadu/venadu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Chera_naduDuke interface {
    feud.Duke
    CKanara伽那罗() 	kanara.KanaraCounty
    CMahoyadapuram摩呼陀耶补罗() 	mahoyadapuram.MahoyadapuramCounty
    CMaldives花环洲() 	maldives.MaldivesCounty
    CQalqut古里() 	qalqut.QalqutCounty
    CVenadu吠那菟() 	venadu.VenaduCounty
}

type 支罗拿都Chera_naduDuke struct {
	feud.BaseDuke
	伽那罗Kanara 	kanara.KanaraCounty
	摩呼陀耶补罗Mahoyadapuram 	mahoyadapuram.MahoyadapuramCounty
	花环洲Maldives 	maldives.MaldivesCounty
	古里Qalqut 	qalqut.QalqutCounty
	吠那菟Venadu 	venadu.VenaduCounty
}

func (d *支罗拿都Chera_naduDuke) CKanara伽那罗() kanara.KanaraCounty {
	return d.伽那罗Kanara
}
    
func (d *支罗拿都Chera_naduDuke) CMahoyadapuram摩呼陀耶补罗() mahoyadapuram.MahoyadapuramCounty {
	return d.摩呼陀耶补罗Mahoyadapuram
}
    
func (d *支罗拿都Chera_naduDuke) CMaldives花环洲() maldives.MaldivesCounty {
	return d.花环洲Maldives
}
    
func (d *支罗拿都Chera_naduDuke) CQalqut古里() qalqut.QalqutCounty {
	return d.古里Qalqut
}
    
func (d *支罗拿都Chera_naduDuke) CVenadu吠那菟() venadu.VenaduCounty {
	return d.吠那菟Venadu
}
    
var DChera_nadu支罗拿都 Chera_naduDuke = &支罗拿都Chera_naduDuke{}

func init() {
	f := DChera_nadu支罗拿都.(*支罗拿都Chera_naduDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "chera_nadu",
		TitleName: "支罗拿都",
		TitleCode: "d_chera_nadu",
		Counties:  map[string]feud.County{},
	}

	f.伽那罗Kanara = kanara.CKanara伽那罗
	f.伽那罗Kanara.SetParent(f)
	
	f.摩呼陀耶补罗Mahoyadapuram = mahoyadapuram.CMahoyadapuram摩呼陀耶补罗
	f.摩呼陀耶补罗Mahoyadapuram.SetParent(f)
	
	f.花环洲Maldives = maldives.CMaldives花环洲
	f.花环洲Maldives.SetParent(f)
	
	f.古里Qalqut = qalqut.CQalqut古里
	f.古里Qalqut.SetParent(f)
	
	f.吠那菟Venadu = venadu.CVenadu吠那菟
	f.吠那菟Venadu.SetParent(f)
	
}
