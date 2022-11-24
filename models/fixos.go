package models

import "encoding/json"

type Fixos struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	Contas      json.Number `json:"contas" gorm:"column:Contas"`
	Assinaturas json.Number `json:"assinaturas" gorm:"column:Assinaturas"`
	Seguros     json.Number `json:"seguros" gorm:"column:Seguros"`
	Mesadas     json.Number `json:"mesadas" gorm:"column:Mesadas"`
	Impostos    json.Number `json:"impostos" gorm:"column:Impostos"`
	Outros      json.Number `json:"outros" gorm:"column:Outros"`
	Mes         json.Number `json:"mes"`
	Ano         json.Number `json:"ano"`
}
