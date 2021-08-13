package main

import (
	postModel "classificator/post-model"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	model := postModel.PostModel{}

	if modelName := os.Getenv("MODEL_FILE"); modelName != "" {
		err := model.TrainFromFile(modelName)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		err := model.FirstTrain()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if predictText := os.Getenv("PREDICT_TEXT"); predictText != "" {
		fmt.Println(model.Predict(predictText))
	}

	if saveName := os.Getenv("SAVE_FILE"); saveName != "" {
		err := model.SaveToFile(saveName)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
