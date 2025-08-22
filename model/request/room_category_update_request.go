package request

type RoomCategoryUpdateRequest struct {
	Id   int64  `validate:"required" json:"id"`
	Name string `validate:"required,max=255,uniqueRoomCategoryUpdate" json:"name"`
}

//func UniqueRoomCategoryUpdate(repo repositories.RoomCategoryRepository, ctx context.Context, tx *sql.Tx, id int64) validator.Func {
//	return func(field validator.FieldLevel) bool {
//		name := field.Field().String()
//
//		roomCategory, err := repo.FindByName(ctx, tx, name)
//		if err != nil {
//			return true
//		}
//
//		if roomCategory != nil {
//			if roomCategory.Id == id {
//				return true
//			} else {
//				return false
//			}
//		}
//
//		return roomCategory == nil
//	}
//}
