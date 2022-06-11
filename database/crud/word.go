package crud

import (
	"entgo.io/ent/dialect/sql"
	"interface_project/ent"
	"interface_project/ent/fileentity"
	"interface_project/ent/word"
)

func (crud Crud) CreateAllWords(bulk []*ent.WordCreate) ([]*ent.Word, error) {
	return crud.Client.Word.CreateBulk(bulk...).Save(*crud.Ctx)
}

func (crud Crud) GetAllWordsByFileID(fileID int) ([]*ent.Word, error) {
	return crud.Client.Word.Query().Where(func(s *sql.Selector) {
		s.Distinct().OnP(s.Select("title").P()).Select("title").Dialect()
	}).All(*crud.Ctx)
}

func (crud Crud) GetAllWordsByTitle(fileID int, title string) ([]*ent.Word, error) {
	return crud.Client.Word.Query().
		Select(word.FieldSentence,
			word.FieldMeaning,
			word.FieldIsPreposition,
			word.FieldDuration,
			word.FieldStart,
			word.FieldEnd,
		).
		Where(word.HasFileWith(fileentity.IDEQ(fileID))).
		Where(word.TitleEQ(title)).
		All(*crud.Ctx)
}
