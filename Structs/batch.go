package main

// import "fmt"

type postingInstructionBatch struct{
	batch_id string
	posting postingInstructions
	batch_details batchDetails
}

type batchDetails struct{
	kope string 
}

type postingInstructions struct{
	fromAccount string
	toAccount string
	amount int32
}


// func main(){
// 	newPosting := postingInstructionBatch{
// 		batch_id: "djfkdj45-45d6fad-asdf545-dsfkl",
// 		posting: postingInstructions{
// 			fromAccount: "IsyLending_new_account_01",
// 			toAccount: "1",
// 			amount: 150,
// 		},
// 		batch_details: batchDetails{
// 			kope: "DKA12122200001478941A",
// 		},
// 	}
// 	fmt.Printf("%+v", newPosting)
// }