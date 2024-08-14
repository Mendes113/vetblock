package repository

import (
	"vetblock/internal/blockchain"

	"gorm.io/gorm"
)

type VetblockRepository struct {
	Db *gorm.DB
}


//implementar errors

func (r *VetblockRepository) CreateVetblockTable() {
	r.Db.AutoMigrate(&blockchain.Block{})
	
}
//save
func (r *VetblockRepository) SaveVetblock(vetblock *blockchain.Block) {
	r.Db.Create(vetblock)
}

func (r *VetblockRepository) FindVetblockById(id int) *blockchain.Block {
	var vetblock blockchain.Block
	r.Db.First(&vetblock, id)
	return &vetblock
}

func (r *VetblockRepository) FindAllVetblocks() []blockchain.Block {
	var vetblocks []blockchain.Block
	r.Db.Find(&vetblocks)
	return vetblocks
}


