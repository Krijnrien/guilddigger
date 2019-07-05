package micro_db
//
//import (
//	"github.com/krijnrien/microguild/pkg/messages"
//	"fmt"
//	"strings"
//)
//
//func (db *MySQLConn) BulkCreateItem(unsavedRows []messages.Item) error {
//	//func BulkInsert(unsavedRows []messages.Item) {
//	valueStrings := make([]string, 0, len(unsavedRows))
//	valueArgs := make([]interface{}, 0, len(unsavedRows)*9)
//	i := 0
//	for _, post := range unsavedRows {
//		i++
//		var placeholderString []int
//		baseI := i * 9
//		placeholderString = append(placeholderString, baseI-8, baseI-7, baseI-6, baseI-5, baseI-4, baseI-3, baseI-2, baseI-1, baseI)
//		valueStrings = append(valueStrings, fmt.Sprintf("($%s)", strings.Trim(strings.Replace(fmt.Sprint(placeholderString), " ", ", $", -1), "[]")))
//		valueArgs = append(valueArgs, post.ID)
//		valueArgs = append(valueArgs, post.Name)
//		valueArgs = append(valueArgs, post.Icon)
//		valueArgs = append(valueArgs, post.Description)
//		valueArgs = append(valueArgs, post.Type)
//		valueArgs = append(valueArgs, post.Rarity)
//		valueArgs = append(valueArgs, post.Level)
//		valueArgs = append(valueArgs, post.VendorValue)
//		valueArgs = append(valueArgs, post.DefaultSkin)
//	}
//	stmt := fmt.Sprintf("INSERT INTO item (item_id, item_name, icon, description, item_type, rarity, item_level, vendorvalue, defaultskin) VALUES %s", strings.Join(valueStrings, ","))
//	_, err := db.Conn.Exec(stmt, valueArgs...)
//	if err != nil {
//		return err
//	}
//	return nil
//}
