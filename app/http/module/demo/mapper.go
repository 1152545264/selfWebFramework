package demo

import demoService "github.com/1152545264/goWebSelf/app/provider/demo"

func UserModelToUserDTOs(models []UserModel) []UserDTO {
	ret := []UserDTO{}
	for _, model := range models {
		t := UserDTO{
			ID:   model.UserId,
			Name: model.Name,
		}
		ret = append(ret, t)
	}
	return ret
}
func StudentsToUserDTOs(students []demoService.Student) []UserDTO {
	ret := []UserDTO{}
	for _, student := range students {
		t := UserDTO{
			ID:   student.ID,
			Name: student.Name,
		}
		ret = append(ret, t)
	}
	return ret
}
