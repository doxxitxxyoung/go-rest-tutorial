package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Drug struct {
//    Id string   //  upper cases ATM
    ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Drugname string
    Smiles string
    Label int
}

type Sample struct {
    ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Drugname string
    Smiles string
    Label int
    Ecfp []int
    Gex []float64
    Dosage float64
    Duration int
    Cellline string
}




/*
    ID primitive.ObjectID
    Drugname string
    Smiles string
    Label int
    id string
    drugname string
    smiles string
    label int
}
*/


