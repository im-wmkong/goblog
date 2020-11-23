package models

import "goblog/pkg/types"

type BaseModel struct {
	ID uint64
}

func (m BaseModel) GetStringID() string {
	return types.Uint64ToString(m.ID)
}
