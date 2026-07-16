package bolejiang

import "fmt"

func (deliver *Deliver) GetPassageRecommendFullPath() string {
	return fmt.Sprintf("%s-%d", deliver.PassageRecommendPath, deliver.PassageRecommendId)
}
