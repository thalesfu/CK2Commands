package maghreb

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaghrebEmpire interface {
	feud.Empire
	KAfrica阿非利加() africa.AfricaKingdom
	KMauretania马格里布() mauretania.MauretaniaKingdom
}

type 马格里布MaghrebEmpire struct {
	feud.BaseEmpire
	阿非利加Africa     africa.AfricaKingdom
	马格里布Mauretania mauretania.MauretaniaKingdom
}

func (e *马格里布MaghrebEmpire) KAfrica阿非利加() africa.AfricaKingdom {
	return e.阿非利加Africa
}

func (e *马格里布MaghrebEmpire) KMauretania马格里布() mauretania.MauretaniaKingdom {
	return e.马格里布Mauretania
}

var EMaghreb马格里布 MaghrebEmpire = &马格里布MaghrebEmpire{}

func init() {
	f := EMaghreb马格里布.(*马格里布MaghrebEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "maghreb",
		TitleName: "马格里布",
		TitleCode: "e_maghreb",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.阿非利加Africa = africa.KAfrica阿非利加
	f.阿非利加Africa.SetParent(f)
	f.马格里布Mauretania = mauretania.KMauretania马格里布
	f.马格里布Mauretania.SetParent(f)
}
