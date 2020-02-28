package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Drug struct {
//    Id string   //  upper cases ATM
    ID primitive.ObjectID
    Drugname string
    Smiles string
    Label int
}
type Sample struct {
    ID primitive.ObjectID
    Drugname string
    Smiles string
    Label string
    ecfp []int
    gex []float64
    dosage float64
    duration int
    cellline string
}




/*
    id string
    drugname string
    smiles string
    label int
}
*/


