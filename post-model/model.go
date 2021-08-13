package post_model

import (
	"encoding/json"
	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/text"
	"os"
)

type PostModel struct {
	groupsNameToId map[string]int
	model          *text.NaiveBayes
}

func (m *PostModel) FirstTrain() error {
	m.groupsNameToId = map[string]int{}
	m.model = train(m, getPosts())
	return nil
}

func (m *PostModel) Predict(str string) string {
	class := m.model.Predict(str)
	for key, value := range m.groupsNameToId {
		if uint8(value) == class {
			return key
		}
	}
	return ""
}

func (m *PostModel) SaveToFile(modelName string) error {
	err := m.model.PersistToFile(modelName + "-model.json")
	if err != nil {
		return err
	}
	str, err := json.Marshal(m.groupsNameToId)
	if err != nil {
		return err
	}
	err = os.WriteFile(modelName+"-map"+".json", str, 0755)
	return err
}

func (m *PostModel) TrainFromFile(modelName string) error {
	m.model = text.NewNaiveBayes(make(chan base.TextDatapoint, 100), uint8(len(m.groupsNameToId)), base.OnlyWords)
	err := m.model.RestoreFromFile(modelName + "-model.json")
	if err != nil {
		return err
	}
	str, err := os.ReadFile(modelName + "-map" + ".json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(str, &m.groupsNameToId)
	return err
}
