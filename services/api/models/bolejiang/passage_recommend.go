package bolejiang

import "fmt"

func (recommend *PassageRecommend) GetFullPath() string {
	return fmt.Sprintf("%s-%d", recommend.Path, recommend.Id)
}
