package crud

import (
	"interface_project/ent"
	"interface_project/ent/fileentity"
	"interface_project/ent/user"
	"os"
	"time"
)

func (crud Crud) CheckFileIsExists(userEmail, filePath, fileName string) bool {

	if fileEntity, err := crud.
		Client.
		FileEntity.
		Query().
		Where(fileentity.HasOwnerWith(user.EmailEqualFold(userEmail))).
		Where(fileentity.PathEqualFold(filePath)).First(*crud.Ctx); err != nil && fileEntity == nil {
		return false
	}
	return true

}

func (crud Crud) AddFileToUser(user *ent.User, file *os.File, path string) (*ent.FileEntity, error) {
	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if createdFile, err := crud.Client.FileEntity.Create().
		SetOwner(user).
		SetCreatedDate(time.Now()).
		SetDeleted(false).
		SetName(file.Name()).
		SetPath(path).
		SetSize(int16(fileStat.Size())).Save(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return createdFile, nil
	}
}

func (crud Crud) GetAllFiles(userEntity *ent.User) ([]*ent.FileEntity, error) {
	return userEntity.QueryFiles().WithWords().All(*crud.Ctx)
}

func (crud Crud) GetFiles(userEntity *ent.User, idList []int) ([]*ent.FileEntity, error) {
	return userEntity.QueryFiles().Where(fileentity.IDIn(idList...)).All(*crud.Ctx)
}

func (crud Crud) DeleteFiles(user *ent.User, idList []int) error {
	return crud.Client.User.UpdateOne(user).RemoveFileIDs(idList...).Exec(*crud.Ctx)
}

func (crud Crud) GetUserFileByID(fileID int, email string) (*ent.FileEntity, error) {
	user, err := crud.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user.QueryFiles().Where(fileentity.IDEQ(fileID)).First(*crud.Ctx)
}
